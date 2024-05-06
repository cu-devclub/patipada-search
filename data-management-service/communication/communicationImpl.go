package communication

import "data-management/rabbitmq"

type CommunicationImpl struct {
	GRPC     GRPCInterface
	RabbitMQ rabbitmq.RabbitMQInterface
}

func NewCommunicationImpl(GRPC GRPCInterface, rabbitMQ rabbitmq.RabbitMQInterface) Communication {
	return &CommunicationImpl{
		GRPC:     GRPC,
		RabbitMQ: rabbitMQ,
	}
}

func (c *CommunicationImpl) Authorization(token string, requiredRole string) (bool, error) {
	return c.GRPC.Authorization(token, requiredRole)
}

func (c *CommunicationImpl) VerifyUsername(username string) (bool, error) {
	return c.GRPC.VerifyUsername(username)
}

func (c *CommunicationImpl) SearchRecord(recordID string) (bool, error) {
	return c.GRPC.SearchRecord(recordID)
}

func (c *CommunicationImpl) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	return c.RabbitMQ.PublishUpdateRecordsToRabbitMQ(payloadName, message)
}
