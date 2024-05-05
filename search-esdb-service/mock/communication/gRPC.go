package mock

import "search-esdb-service/communication"

func NewMockgRPC() communication.GRPCInterface {
	mockMlGatewayService := NewMockMlGateayClient()
	return &communication.GRPCStruct{
		MlGatewayClient: mockMlGatewayService,
	}
}

