package oidc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	capoidc "github.com/hashicorp/cap/oidc"
	"go.flipt.io/flipt/errors"
	"go.flipt.io/flipt/internal/config"
	"go.flipt.io/flipt/internal/server/auth/method"
	storageauth "go.flipt.io/flipt/internal/storage/auth"
	"go.flipt.io/flipt/rpc/flipt/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	storageMetadataOIDCProvider          = "io.flipt.auth.oidc.provider"
	storageMetadataOIDCEmail             = "io.flipt.auth.oidc.email"
	storageMetadataOIDCEmailVerified     = "io.flipt.auth.oidc.email_verified"
	storageMetadataOIDCName              = "io.flipt.auth.oidc.name"
	storageMetadataOIDCProfile           = "io.flipt.auth.oidc.profile"
	storageMetadataOIDCPicture           = "io.flipt.auth.oidc.picture"
	storageMetadataOIDCSub               = "io.flipt.auth.oidc.sub"
	storageMetadataOIDCPreferredUsername = "io.flipt.auth.oidc.preferred_username"
)

// errProviderNotFound is returned when a provider is requested which
// was not configured
var errProviderNotFound = errors.ErrNotFound("provider not found")

// PCKEVerifier is a code verifier used for a PKCE flow during OIDC authentication.
// This value is declared outside the scope of the function because of consistency throughout
// the authenciation legs of OIDC.
var PKCEVerifier, _ = capoidc.NewCodeVerifier()

// Server is the core OIDC server implementation for Flipt.
// It supports two primary operations:
// - AuthorizeURL
// - Callback
// These are two legs of the OIDC/OAuth flow.
// Step 1 is Flipt establishes a URL directed at the delegated authentication service (e.g. Google).
// The URL is configured using the client ID configured for the provided, a state parameter used to
// prevent CSRF attacks and a callback URL directing back to the Callback operation.
// Step 2 the user-agent navigates to the authorizer and establishes authenticity with them.
// Once established they're redirected to the Callback operation with an authenticity code.
// Step 3 the Callback operation uses this "code" and exchanges with the authorization service
// for an ID Token. The validity of the response is checked (signature verified) and then the identity
// details contained in this response are used to create a temporary Flipt client token.
// This client token can be used to access the rest of the Flipt API.
// Given the user-agent is requestin using HTTP the token is instead established as an HTTP cookie.
type Server struct {
	logger *zap.Logger
	store  storageauth.Store
	config config.AuthenticationConfig

	auth.UnimplementedAuthenticationMethodOIDCServiceServer
}

func NewServer(
	logger *zap.Logger,
	store storageauth.Store,
	config config.AuthenticationConfig,
) *Server {
	return &Server{
		logger: logger,
		store:  store,
		config: config,
	}
}

// RegisterGRPC registers the server as an Server on the provided grpc server.
func (s *Server) RegisterGRPC(server *grpc.Server) {
	auth.RegisterAuthenticationMethodOIDCServiceServer(server, s)
}

func (s *Server) SkipsAuthentication(ctx context.Context) bool {
	return true
}

// AuthorizeURL constructs and returns a URL directed at the requested OIDC provider
// based on our internal oauth2 client configuration.
// The operation is configured to return a URL which ultimately redirects to the
// callback operation below.
func (s *Server) AuthorizeURL(ctx context.Context, req *auth.AuthorizeURLRequest) (*auth.AuthorizeURLResponse, error) {
	provider, oidcRequest, err := s.providerFor(req.Provider, req.State)
	if err != nil {
		return nil, fmt.Errorf("authorize: %w", err)
	}

	// Create an auth URL
	authURL, err := provider.AuthURL(context.Background(), oidcRequest)
	if err != nil {
		return nil, err
	}

	return &auth.AuthorizeURLResponse{AuthorizeUrl: authURL}, nil
}

// Callback attempts to authenticate a callback request from a delegated authorization service.
// Given the request includes a "state" parameter then the requests metadata is interrogated
// for the "flipt_client_state" metadata key.
// This entry must exist and the value match the request state.
// The provided code is exchanged with the associated authorization service provider for an "id_token".
// We verify the retrieved "id_token" is valid and for our client.
// Once verified we extract the users associated email address.
// Given all this completes successfully then we established an associated clientToken in
// the backing authentication store with the identity information retrieved as metadata.
func (s *Server) Callback(ctx context.Context, req *auth.CallbackRequest) (_ *auth.CallbackResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("handling OIDC callback: %w", err)
		}
	}()

	if req.State != "" {
		if err := method.CallbackValidateState(ctx, req.State); err != nil {
			return nil, err
		}
	}

	provider, oidcRequest, err := s.providerFor(req.Provider, req.State)
	if err != nil {
		return nil, err
	}

	responseToken, err := provider.Exchange(ctx, oidcRequest, req.State, req.Code)
	if err != nil {
		return nil, err
	}

	metadata := map[string]string{
		storageMetadataOIDCProvider: req.Provider,
	}

	// Extract custom claims
	var claims claims
	if err := responseToken.IDToken().Claims(&claims); err != nil {
		return nil, err
	}
	claims.addToMetadata(metadata)

	clientToken, a, err := s.store.CreateAuthentication(ctx, &storageauth.CreateAuthenticationRequest{
		Method:    auth.Method_METHOD_OIDC,
		ExpiresAt: timestamppb.New(time.Now().UTC().Add(s.config.Session.TokenLifetime)),
		Metadata:  metadata,
	})
	if err != nil {
		return nil, err
	}

	return &auth.CallbackResponse{
		ClientToken:    clientToken,
		Authentication: a,
	}, nil
}

func callbackURL(host, provider string) string {
	// strip trailing slash from host
	host = strings.TrimSuffix(host, "/")
	return host + "/auth/v1/method/oidc/" + provider + "/callback"
}

func (s *Server) providerFor(provider string, state string) (*capoidc.Provider, *capoidc.Req, error) {
	var (
		config   *capoidc.Config
		callback string
	)

	pConfig, ok := s.config.Methods.OIDC.Method.Providers[provider]
	if !ok {
		return nil, nil, fmt.Errorf("requested provider %q: %w", provider, errProviderNotFound)
	}

	callback = callbackURL(pConfig.RedirectAddress, provider)

	var err error
	config, err = capoidc.NewConfig(
		pConfig.IssuerURL,
		pConfig.ClientID,
		capoidc.ClientSecret(pConfig.ClientSecret),
		[]capoidc.Alg{oidc.RS256},
		[]string{callback},
	)
	if err != nil {
		return nil, nil, err
	}

	p, err := capoidc.NewProvider(config)
	if err != nil {
		return nil, nil, err
	}

	var oidcOpts = []capoidc.Option{
		capoidc.WithState(state),
		capoidc.WithScopes(pConfig.Scopes...),
		capoidc.WithNonce("static"), // TODO(georgemac): dropping nonce for now
	}

	if pConfig.UsePKCE {
		oidcOpts = append(oidcOpts, capoidc.WithPKCE(PKCEVerifier))
	}

	req, err := capoidc.NewRequest(2*time.Minute, callback,
		oidcOpts...,
	)
	if err != nil {
		return nil, nil, err
	}

	return p, req, nil
}

type claims struct {
	Email    *string `json:"email"`
	Verified *bool   `json:"email_verified"`
	Name     *string `json:"name"`
	Profile  *string `json:"profile"`
	Picture  *string `json:"picture"`
	Sub      *string `json:"sub"`
}

func (c claims) addToMetadata(m map[string]string) {
	set := func(key string, s *string) {
		if s != nil && *s != "" {
			m[key] = *s
		}
	}

	set(storageMetadataOIDCEmail, c.Email)
	set(storageMetadataOIDCName, c.Name)
	set(storageMetadataOIDCProfile, c.Profile)
	set(storageMetadataOIDCPicture, c.Picture)
	set(storageMetadataOIDCSub, c.Sub)

	if c.Verified != nil {
		m[storageMetadataOIDCEmailVerified] = fmt.Sprintf("%v", *c.Verified)
	}
}
