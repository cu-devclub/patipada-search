package rabbitmq

import (
	"encoding/json"
	"fmt"
	"search-esdb-service/config"
	"search-esdb-service/constant"
	recordUsecases "search-esdb-service/record/usecases"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn          *amqp.Connection
	recordUsecase recordUsecases.RecordUsecase
}

func NewConsumer(conn *amqp.Connection, cfg *config.Config, recordUsecase recordUsecases.RecordUsecase) (Consumer, error) {
	consumer := Consumer{
		conn:          conn,
		recordUsecase: recordUsecase,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel, constant.RECORD_EXCHANGE)
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			constant.RECORD_EXCHANGE,
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)
			go consumer.handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [topic, %s]\n", q.Name)
	<-forever

	return nil
}
