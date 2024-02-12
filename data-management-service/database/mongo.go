package database

import (
	"context"
	"data-management/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	Db *mongo.Client
}

func NewMongoDatabase(cfg *config.Config) Database {
	// create a connection to mongo db
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &mongoDatabase{
		Db: client,
	}


}

func (m *mongoDatabase) GetDb() *mongo.Client {
	return m.Db
}
