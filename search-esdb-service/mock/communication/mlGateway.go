package mock

import (
	"context"
	"search-esdb-service/proto/ml_gateway_proto"

	"google.golang.org/grpc"
)

type MockMlGatewayClientInterface interface {
	Text2Vec(ctx context.Context, in *ml_gateway_proto.Text2VecRequest, opts ...grpc.CallOption) (*ml_gateway_proto.Text2VecResponse, error)
}

type MockMlGatewayClient struct {
	mlGatewayClient MockMlGatewayClientInterface
}

func NewMockMlGateayClient() MockMlGatewayClientInterface {
	return &MockMlGatewayClient{}
}

func MockText2VecResponse() *ml_gateway_proto.Text2VecResponse {
	res := &ml_gateway_proto.Result{
		Name:      "mock",
		Embedding: []float32{0.1, 0.2, 0.3},
		Score:     0.5,
	}
	return &ml_gateway_proto.Text2VecResponse{
		Results: []*ml_gateway_proto.Result{res},
	}
}

func (g *MockMlGatewayClient) Text2Vec(ctx context.Context, in *ml_gateway_proto.Text2VecRequest, opts ...grpc.CallOption) (*ml_gateway_proto.Text2VecResponse, error) {
	return MockText2VecResponse(), nil
}
