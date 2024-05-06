package mock_communication

import (
	"data-management/communication"
	"data-management/rabbitmq"
)

type MockCommunicationStruct struct {
	GRPC communication.GRPCInterface
	RabbitMQ rabbitmq.RabbitMQInterface
}

func MockCommunication() communication.Communication {
	gRPC := NewMockgRPC()
	rabbitMQ := MockRabbitMQ()

	return MockCommunicationStruct{
		GRPC:     gRPC,
		RabbitMQ: rabbitMQ,
	}
}

func (m MockCommunicationStruct) Authorization(token string, requiredRole string) (bool, error) {
	return m.GRPC.Authorization(token, requiredRole)
}

func (m MockCommunicationStruct) VerifyUsername(username string) (bool, error) {
	return m.GRPC.VerifyUsername(username)
}

func (m MockCommunicationStruct) SearchRecord(recordID string) (bool, error) {
	return m.GRPC.SearchRecord(recordID)
}

func (m MockCommunicationStruct) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	return m.RabbitMQ.PublishUpdateRecordsToRabbitMQ(payloadName, message)
}



