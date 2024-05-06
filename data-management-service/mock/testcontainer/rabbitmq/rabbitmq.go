package test_container_rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
)

func MockRabbitMQ() (*amqp.Connection, func()) {
	// connect to rabbit mq
	ctx := context.Background()

	rabbitmqContainer, err := rabbitmq.RunContainer(ctx,
		testcontainers.WithImage("rabbitmq:3.12.11-management-alpine"),
		rabbitmq.WithAdminPassword("testCOntinaerPassword"),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	host, err := rabbitmqContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}

	port, err := rabbitmqContainer.MappedPort(ctx, "5672")
	if err != nil {
		log.Fatalf("failed to get port: %s", err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@" + host + ":" + port.Port())
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err)
	}

	return conn, func() {
		rabbitmqContainer.Terminate(ctx)
	}
}
