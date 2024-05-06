package mock_communication

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

var searchResponse *bool

func SetSearchResponse(response bool) {
	searchResponse = &response
}

func MockVSearchRecordResponse() *search_proto.SearchResponse {
	return &search_proto.SearchResponse{
		IsFounded: *searchResponse, // or false, depending on what you want to mock
	}
}

func (c *MockSearchClient) SearchRecord(ctx context.Context, in *search_proto.SearchRequest, opts ...grpc.CallOption) (*search_proto.SearchResponse, error) {
	return MockVSearchRecordResponse(), nil
}
