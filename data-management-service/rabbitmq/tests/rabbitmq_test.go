package rabbitmq_test

import (
	"data-management/constant"
	test_container_rabbitmq "data-management/mock/testcontainer/rabbitmq"
	rabbit "data-management/rabbitmq"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	amqp "github.com/rabbitmq/amqp091-go"
)

func setupRabbitMQ() (rabbit.RabbitMQInterface, func()) {
	conn, cleanup := test_container_rabbitmq.MockRabbitMQ()

	emit, err := rabbit.NewEmitter(conn)
	if err != nil {
		panic(err)
	}

	return &rabbit.RabbitMQStruct{
		Conn:    conn,
		Emitter: emit,
	}, cleanup
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}

func setupConsumer(client rabbit.RabbitMQInterface, queueName string) (<-chan amqp.Delivery, error) {
	conn := client.GetConnection()
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err = rabbit.DeclareExchange(ch, constant.RECORD_EXCHANGE); err != nil {
		return nil, err
	}

	q, err := declareQueue(ch)
	if err != nil {
		return nil, err
	}

	if err = ch.QueueBind(q.Name, queueName, constant.RECORD_EXCHANGE, false, nil); err != nil {
		return nil, err
	}

	return ch.Consume(q.Name, "", true, false, false, false, nil)
}

func TestRabbitMQ(t *testing.T) {
	client, cleanup := setupRabbitMQ()
	defer cleanup()

	t.Run("Test emit message", func(t *testing.T) {
		msgs, err := setupConsumer(client, "test")
		if err != nil {
			t.Fatal(err)
		}
		emitter := client.GetEmitter()
		expectedMessage := "Hello, World!"
		assert.NoError(t, emitter.Emit(expectedMessage, "test"))

		select {
		case msg := <-msgs:
			assert.Equal(t, expectedMessage, string(msg.Body))
		case <-time.After(5 * time.Second):
			t.Fatal("did not receive message")
		}
	})

	t.Run("Test publish update records to RabbitMQ", func(t *testing.T) {
		msgs, err := setupConsumer(client, constant.UPDATE_RECORD_TOPIC)
		if err != nil {
			t.Fatal(err)
		}

		payloadName := "payload_name"
		message := "message"
		assert.NoError(t, client.PublishUpdateRecordsToRabbitMQ(payloadName, message))

		expectedRes := rabbit.RabbitMQPayload{
			Name: payloadName,
			Data: message,
		}
		expectedJSON, _ := json.Marshal(expectedRes)

		select {
		case msg := <-msgs:
			assert.Equal(t, string(expectedJSON), string(msg.Body))
		case <-time.After(5 * time.Second):
			t.Fatal("did not receive message")
		}
	})
}
