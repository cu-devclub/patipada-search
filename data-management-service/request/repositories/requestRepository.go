package repositories

import (
	"data-management/communication"

	"go.mongodb.org/mongo-driver/mongo"
)

type requestRepositories struct {
	mongo                    *mongo.Client
	requestCollection        *mongo.Collection
	requestCounterCollection *mongo.Collection
	communicationClient      communication.Communication
}

func NewRequestRepositories(mongo *mongo.Client,c communication.Communication) Repositories {
	requestCollection := mongo.Database("request").Collection("request")
	requestCounterCollection := mongo.Database("request").Collection("counters")
	return &requestRepositories{
		mongo:                    mongo,
		requestCollection:        requestCollection,
		requestCounterCollection: requestCounterCollection,
		communicationClient:      c,
	}
}
