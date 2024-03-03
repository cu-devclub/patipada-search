package communication

import (
	"fmt"
	"log"
	"math"
	"search-esdb-service/config"
	"search-esdb-service/event"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	amqp "github.com/rabbitmq/amqp091-go"

	recordESRepositories "search-esdb-service/record/repositories/recordRepository"
	recordUsecases "search-esdb-service/record/usecases"
)

type RabbitMQStruct struct {
	Conn     *amqp.Connection
	Consumer event.Consumer
}

func ConnectToRabbitMQ(cfg *config.Config, db *elasticsearch.Client) (*RabbitMQStruct, error) {
	log.Println("Connecting to RabbitMQ...")
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	connectionURL := fmt.Sprintf("amqp://%s:%s@%s/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.URL,
	)
	log.Println("Rabbit MQ connection URL:", connectionURL)
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

	log.Println("Connected to RabbitMQ!")

	recordRepository := recordESRepositories.NewRecordESRepository(db)
	recordUsecases := recordUsecases.NewRecordUsecase(recordRepository, nil)

	consumer, err := event.NewConsumer(connection, cfg, recordUsecases)
	if err != nil {
		return nil, err
	}

	return &RabbitMQStruct{
		Conn:     connection,
		Consumer: consumer,
	}, nil
}

func (c *CommunicationImpl) Listen(topics []string) error {
	log.Println("Listening to topics...", topics)
	return c.RabbitMQ.Consumer.Listen(topics)
}
