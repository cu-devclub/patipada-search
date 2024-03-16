package communication

import (
	"fmt"
	"log/slog"
	"math"
	"search-esdb-service/config"
	"search-esdb-service/event"
	"search-esdb-service/server"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQStruct struct {
	Conn     *amqp.Connection
	Consumer event.Consumer
}

func ConnectToRabbitMQ(cfg *config.Config, db *elasticsearch.Client, recordArch server.RecordArch) (*RabbitMQStruct, error) {
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
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		time.Sleep(backOff)
	}

	consumer, err := event.NewConsumer(connection, cfg, recordArch.Usecase)
	if err != nil {
		return nil, err
	}
	slog.Info("Creating Consumer successfully!")

	return &RabbitMQStruct{
		Conn:     connection,
		Consumer: consumer,
	}, nil
}

func (c *CommunicationImpl) Listen(topics []string) error {
	slog.Info("Listening to", slog.Any("topics", topics))
	return c.RabbitMQ.Consumer.Listen(topics)
}
