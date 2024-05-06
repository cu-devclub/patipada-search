package test_container_database

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMockMongoClient() (*mongo.Client, func(), error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.4.4-bionic",
		ExposedPorts: []string{"27017/tcp", "27018/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "root",
			"MONGO_INITDB_ROOT_PASSWORD": "example",
		},
	}

	mongodbC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("error starting mongodb container: %s", err)
	}

	host, _ := mongodbC.Host(ctx)
	p, _ := mongodbC.MappedPort(ctx, "27017/tcp")
	port := p.Int()

	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/",
		"root", "example", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("cannot connect: %v\n", err)
		return nil, nil, nil
	}

	return client, func() {
		mongodbC.Terminate(ctx)
	}, nil
}
