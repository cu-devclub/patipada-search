package mock_communication

import (
	"data-management/config"
	"data-management/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MockEmitterStruct struct {
	conn *amqp.Connection
}

func MockEmitter() rabbitmq.EmitterInterface {
	return &MockEmitterStruct{
		conn: nil,
	}
}

func (e *MockEmitterStruct) Emit(message string, key string) error {
	return nil
}

type MockRabbitMQStruct struct {
	Conn         *amqp.Connection
	Emitter      rabbitmq.EmitterInterface
	Rabbitconfig *config.RabbitMQ
}

func MockRabbitMQ() rabbitmq.RabbitMQInterface {
	// Get the emitter
	emitter := MockEmitter()

	return &MockRabbitMQStruct{
		Conn:         nil,
		Emitter:      emitter,
		Rabbitconfig: nil,
	}
}

func (r *MockRabbitMQStruct) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	return nil
}

func (r *MockRabbitMQStruct) GetConnection() *amqp.Connection {
	return r.Conn
}

func (r *MockRabbitMQStruct) GetEmitter() rabbitmq.EmitterInterface {
	return r.Emitter
}

func (r *MockRabbitMQStruct) CloseConnection() {
	return
}
