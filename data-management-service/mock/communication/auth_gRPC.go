package mock_communication

import (
	"context"
	"data-management/proto/auth_proto"

	"google.golang.org/grpc"
)

var authorizeResponse *bool

var verifyUsernameResponse *bool

func init() {
	b := true
	authorizeResponse = &b
	verifyUsernameResponse = &b
}

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


func SetAuthorizationResponse(response bool) {
	authorizeResponse = &response
}

func MockAuthorizationResponse() *auth_proto.AuthorizationResponse {
	return &auth_proto.AuthorizationResponse{
		IsAuthorized: *authorizeResponse, 
	}
}

func (c *MockAuthServiceClient) Authorization(ctx context.Context, in *auth_proto.AuthorizationRequest, opts ...grpc.CallOption) (*auth_proto.AuthorizationResponse, error) {
	return MockAuthorizationResponse(), nil
}


func SetVerifyUsernameResponse(response bool) {
	verifyUsernameResponse = &response
}

func MockVerifyUsernameResponse() *auth_proto.VerifyUsernameResponse {
	return &auth_proto.VerifyUsernameResponse{
		IsVerified: *verifyUsernameResponse, 
	}
}

func (c *MockAuthServiceClient) VerifyUsername(ctx context.Context, in *auth_proto.VerifyUsernameRequest, opts ...grpc.CallOption) (*auth_proto.VerifyUsernameResponse, error) {
	return MockVerifyUsernameResponse(), nil
}
