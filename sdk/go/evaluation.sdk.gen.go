// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.

package sdk

import (
	context "context"
	evaluation "go.flipt.io/flipt/rpc/flipt/evaluation"
)

type EvaluationClient interface {
	EvaluationServiceClient() evaluation.EvaluationServiceClient
}

type Evaluation struct {
	transport     EvaluationClient
	tokenProvider ClientTokenProvider
}

type EvaluationService struct {
	transport     evaluation.EvaluationServiceClient
	tokenProvider ClientTokenProvider
}

func (s Evaluation) EvaluationService() *EvaluationService {
	return &EvaluationService{
		transport:     s.transport.EvaluationServiceClient(),
		tokenProvider: s.tokenProvider,
	}
}
func (x *EvaluationService) Boolean(ctx context.Context, v *evaluation.EvaluationRequest) (*evaluation.BooleanEvaluationResponse, error) {
	ctx, err := authenticate(ctx, x.tokenProvider)
	if err != nil {
		return nil, err
	}
	return x.transport.Boolean(ctx, v)
}

func (x *EvaluationService) Variant(ctx context.Context, v *evaluation.EvaluationRequest) (*evaluation.VariantEvaluationResponse, error) {
	ctx, err := authenticate(ctx, x.tokenProvider)
	if err != nil {
		return nil, err
	}
	return x.transport.Variant(ctx, v)
}

func (x *EvaluationService) Batch(ctx context.Context, v *evaluation.BatchEvaluationRequest) (*evaluation.BatchEvaluationResponse, error) {
	ctx, err := authenticate(ctx, x.tokenProvider)
	if err != nil {
		return nil, err
	}
	return x.transport.Batch(ctx, v)
}
