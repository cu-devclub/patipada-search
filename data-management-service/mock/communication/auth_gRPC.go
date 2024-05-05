package mock

import (
	"context"
	"data-management/proto/auth_proto"

	"google.golang.org/grpc"
)

type MockAuthServiceClientInterface interface {
	Authorization(ctx context.Context, in *auth_proto.AuthorizationRequest, opts ...grpc.CallOption) (*auth_proto.AuthorizationResponse, error)
	VerifyUsername(ctx context.Context, in *auth_proto.VerifyUsernameRequest, opts ...grpc.CallOption) (*auth_proto.VerifyUsernameResponse, error)
}

type MockAuthServiceClient struct {
	authServiceClient MockAuthServiceClientInterface
}

func NewMockAuthServiceClient() MockAuthServiceClientInterface {
	return &MockAuthServiceClient{}
}

func MockAuthorizationResponse() *auth_proto.AuthorizationResponse {
	return &auth_proto.AuthorizationResponse{
		IsAuthorized: true, // or false, depending on what you want to mock
	}
}

func (c *MockAuthServiceClient) Authorization(ctx context.Context, in *auth_proto.AuthorizationRequest, opts ...grpc.CallOption) (*auth_proto.AuthorizationResponse, error) {
	return MockAuthorizationResponse(), nil
}

func MockVerifyUsernameResponse() *auth_proto.VerifyUsernameResponse {
	return &auth_proto.VerifyUsernameResponse{
		IsVerified: true, // or false, depending on what you want to mock
	}
}

func (c *MockAuthServiceClient) VerifyUsername(ctx context.Context, in *auth_proto.VerifyUsernameRequest, opts ...grpc.CallOption) (*auth_proto.VerifyUsernameResponse, error) {
	return MockVerifyUsernameResponse(), nil
}
