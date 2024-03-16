package communication

type CommunicationImpl struct {
	GRPC GRPCStruct
	RabbitMQ RabbitMQStruct
	
}

func NewCommunicationImpl(GRPC GRPCStruct,rabbitMQ RabbitMQStruct) Communication {
	return &CommunicationImpl{
		GRPC: GRPC,
		RabbitMQ: rabbitMQ,
	}
}