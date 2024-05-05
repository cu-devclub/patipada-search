package mock

import (
	"context"
	"data-management/proto/search_proto"

	"google.golang.org/grpc"
)

type MockSearchServiceClientInterface interface {
	SearchRecord(ctx context.Context, in *search_proto.SearchRequest, opts ...grpc.CallOption) (*search_proto.SearchResponse, error)
}

type MockSearchClient struct {
	searchClient MockSearchServiceClientInterface
}

func NewMockSearchServiceClient() MockSearchServiceClientInterface {
	return &MockSearchClient{}
}

func MockVSearchRecordResponse() *search_proto.SearchResponse {
	return &search_proto.SearchResponse{
		IsFounded: true, // or false, depending on what you want to mock
	}
}

func (c *MockSearchClient) SearchRecord(ctx context.Context, in *search_proto.SearchRequest, opts ...grpc.CallOption) (*search_proto.SearchResponse, error) {
	return MockVSearchRecordResponse(), nil
}
