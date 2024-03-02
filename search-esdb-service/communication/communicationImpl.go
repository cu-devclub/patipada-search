package communication

type CommunicationImpl struct {
	RabbitMQ *RabbitMQStruct
	
}

func NewCommunicationImpl(rabbitMQ RabbitMQStruct) Communication {
	return &CommunicationImpl{
		RabbitMQ: &rabbitMQ,
	}
}