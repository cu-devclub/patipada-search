package mock

import "data-management/communication"

func NewMockgRPC() communication.GRPCInterface {
	mockAuthService := NewMockAuthServiceClient()
	mockSearchService := NewMockSearchServiceClient()
	return &communication.GRPCStruct{
		AuthClient:   mockAuthService,
		SearchClient: mockSearchService,
	}
}
