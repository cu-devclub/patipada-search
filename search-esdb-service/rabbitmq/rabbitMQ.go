package rabbitmq

import (
	"fmt"
	"log/slog"
	"math"
	"search-esdb-service/config"
	recordUsecase "search-esdb-service/record/usecases"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQStruct struct {
	Conn     *amqp.Connection
	Consumer Consumer
}

func ConnectToRabbitMQ(cfg *config.Config, db *elasticsearch.Client, recordUsecase recordUsecase.RecordUsecase) (*RabbitMQStruct, error) {
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

	consumer, err := NewConsumer(connection, cfg, recordUsecase)
	if err != nil {
		return nil, err
	}
	slog.Info("Creating Consumer successfully!")

	return &RabbitMQStruct{
		Conn:     connection,
		Consumer: consumer,
	}, nil
}

func (c *RabbitMQStruct) Listen(topics []string) error {
	slog.Info("Listening to", slog.Any("topics", topics))
	return c.Consumer.Listen(topics)
}
