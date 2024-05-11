package rabbitmq

import (
	"data-management/config"
	"data-management/constant"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQInterface interface {
	PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error
	GetEmitter() EmitterInterface
	GetConnection() *amqp.Connection
	CloseConnection()
}

type RabbitMQStruct struct {
	Conn    *amqp.Connection
	Emitter EmitterInterface
}

func ConnectToRabbitMQ(cfg *config.Config) (RabbitMQInterface, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	connectionURL := fmt.Sprintf("amqp://%s:%s@%s/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.URL,
	)
	// don't continue until we have a connection
	for {
		c, err := amqp.Dial(connectionURL)
		if err != nil {
			slog.Warn("Rabbit MQ is not ready yet....")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			slog.Warn("Rabbit MQ is not ready, giving up....", err)
			return nil, nil
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		time.Sleep(backOff)
	}

	// Get the emitter
	emitter, err := NewEmitter(connection)
	if err != nil {
		return nil, err
	}

	return &RabbitMQStruct{
		Conn:    connection,
		Emitter: emitter,
	}, nil
}

type RabbitMQPayload struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (c *RabbitMQStruct) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	// convert message to payload json string
	payload := RabbitMQPayload{
		Name: payloadName,
		Data: message,
	}
	j, _ := json.Marshal(payload)
	err := c.Emitter.Emit(string(j), constant.UPDATE_RECORD_TOPIC)
	if err != nil {
		return err
	}
	return nil
}

func (c *RabbitMQStruct) GetConnection() *amqp.Connection {
	return c.Conn
}

func (c *RabbitMQStruct) GetEmitter() EmitterInterface {
	return c.Emitter
}

func (c *RabbitMQStruct) CloseConnection() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}
