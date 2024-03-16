package communication

import (
	"data-management/config"
	"data-management/constant"
	"data-management/event"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQStruct struct {
	Conn         *amqp.Connection
	Emitter      *event.Emitter
	rabbitconfig *config.RabbitMQ
}

func ConnectToRabbitMQ(cfg *config.Config) (*RabbitMQStruct, error) {
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
	emitter, err := event.NewEmitter(connection, cfg)
	if err != nil {
		return nil, err
	}

	return &RabbitMQStruct{
		Conn:         connection,
		Emitter:      emitter,
		rabbitconfig: &cfg.RabbitMQ,
	}, nil
}

type RabbitMQPayload struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (c *CommunicationImpl) PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error {
	// convert message to payload json string
	payload := RabbitMQPayload{
		Name: payloadName,
		Data: message,
	}
	j, _ := json.Marshal(payload)
	err := c.RabbitMQ.Emitter.Emit(string(j), constant.UPDATE_RECORD_TOPIC)
	if err != nil {
		return err
	}
	return nil
}
