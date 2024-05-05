package mock

import (
	"data-management/communication"
	"data-management/config"
	"data-management/event"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MockRabbitMQStruct struct {
	Conn         *amqp.Connection
	Emitter      *event.Emitter
	Rabbitconfig *config.RabbitMQ
}

func MockRabbitMQ() communication.RabbitMQInterface {
	return &MockRabbitMQStruct{
		Conn:         nil,
		Emitter:      nil,
		Rabbitconfig: nil,
	}
}

func (r *MockRabbitMQStruct) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	return nil
}


func (r *MockRabbitMQStruct) CloseConnection() {
	return
}