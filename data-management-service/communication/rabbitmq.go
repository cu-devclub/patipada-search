package communication

import (
	"data-management/config"
	"data-management/constant"
	"data-management/event"
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("Connecting to Rabbit MQ....")
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	connectionURL := fmt.Sprintf("amqp://%s:%s@%s/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.URL,
	)
	log.Println("Connecting to Rabbit MQ with connection URL:", connectionURL)
	// don't continue until we have a connection
	for {
		c, err := amqp.Dial(connectionURL)
		if err != nil {
			fmt.Println("Rabbit MQ is not ready yet....")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println("Rabbit MQ is not ready, giving up....", err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		time.Sleep(backOff)
	}

	log.Println("Connected to Rabbit MQ!")

	// Get the emitter
	emitter, err := event.NewEmitter(connection, cfg)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully initialized RabbitMQ connection & emitter!")
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
	log.Println("Publish update records to RabbitMQ with message", message)
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
	log.Println("Published to RabbitMQ!!!!")
	return nil
}
