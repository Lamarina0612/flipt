// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.

package sdk

import (
	context "context"
	flipt "go.flipt.io/flipt/rpc/flipt"
	data "go.flipt.io/flipt/rpc/flipt/data"
	evaluation "go.flipt.io/flipt/rpc/flipt/evaluation"
	meta "go.flipt.io/flipt/rpc/flipt/meta"
	metadata "google.golang.org/grpc/metadata"
)

type Transport interface {
	AuthClient() AuthClient
	FliptClient() flipt.FliptClient
	DataClient() data.DataServiceClient
	EvaluationClient() evaluation.EvaluationServiceClient
	MetaClient() meta.MetadataServiceClient
}

// ClientTokenProvider is a type which when requested provides a
// client token which can be used to authenticate RPC/API calls
// invoked through the SDK.
type ClientTokenProvider interface {
	ClientToken() (string, error)
}

// SDK is the definition of Flipt's Go SDK.
// It depends on a pluggable transport implementation and exposes
// a consistent API surface area across both transport implementations.
// It also provides consistent client-side instrumentation and authentication
// lifecycle support.
type SDK struct {
	transport     Transport
	tokenProvider ClientTokenProvider
}

// Option is a functional option which configures the Flipt SDK.
type Option func(*SDK)

// WithClientTokenProviders returns an Option which configures
// any supplied SDK with the provided ClientTokenProvider.
func WithClientTokenProvider(p ClientTokenProvider) Option {
	return func(s *SDK) {
		s.tokenProvider = p
	}
}

// StaticClientTokenProvider is a string which is supplied as a static client token
// on each RPC which requires authentication.
type StaticClientTokenProvider string

// ClientToken returns the underlying string that is the StaticClientTokenProvider.
func (p StaticClientTokenProvider) ClientToken() (string, error) {
	return string(p), nil
}

// New constructs and configures a Flipt SDK instance from
// the provided Transport implementation and options.
func New(t Transport, opts ...Option) SDK {
	sdk := SDK{transport: t}

	for _, opt := range opts {
		opt(&sdk)
	}

	return sdk
}

func (s SDK) Auth() *Auth {
	return &Auth{
		transport:     s.transport.AuthClient(),
		tokenProvider: s.tokenProvider,
	}
}

func (s SDK) Flipt() *Flipt {
	return &Flipt{
		transport:     s.transport.FliptClient(),
		tokenProvider: s.tokenProvider,
	}
}

func (s SDK) Data() *Data {
	return &Data{
		transport:     s.transport.DataClient(),
		tokenProvider: s.tokenProvider,
	}
}

func (s SDK) Evaluation() *Evaluation {
	return &Evaluation{
		transport:     s.transport.EvaluationClient(),
		tokenProvider: s.tokenProvider,
	}
}

func (s SDK) Meta() *Meta {
	return &Meta{
		transport:     s.transport.MetaClient(),
		tokenProvider: s.tokenProvider,
	}
}

func authenticate(ctx context.Context, p ClientTokenProvider) (context.Context, error) {
	if p != nil {
		token, err := p.ClientToken()
		if err != nil {
			return ctx, err
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	}

	return ctx, nil
}
