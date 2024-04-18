package event

import (
	"context"
	"data-management/config"
	"data-management/constant"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	conn *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return declareExchange(channel, constant.RECORD_EXCHANGE)
}

func NewEmitter(conn *amqp.Connection, cfg *config.Config) (*Emitter, error) {
	log.Println("Creating new emitter.....")
	emitter := &Emitter{
		conn: conn,
	}

	err := emitter.setup()
	if err != nil {
		return &Emitter{}, err
	}

	log.Println("Success creating new emitter.....")
	return emitter, nil
}

func (e *Emitter) Emit(message string, key string) error {
	// key == topics in RabbitMQ
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = channel.PublishWithContext(
		ctx,
		constant.RECORD_EXCHANGE,
		key, // topics == key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
