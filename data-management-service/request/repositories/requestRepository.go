package repositories

import "go.mongodb.org/mongo-driver/mongo"

type requestRepositories struct {
	mongo                   *mongo.Client
	requestCollection       *mongo.Collection
	requestCounterCollection *mongo.Collection
}

func NewRequestRepositories(mongo *mongo.Client) Repositories {
	requestCollection := mongo.Database("request").Collection("request")
	requestCounterCollection := mongo.Database("request").Collection("counters")
	return &requestRepositories{
		mongo:                    mongo,
		requestCollection:        requestCollection,
		requestCounterCollection: requestCounterCollection,
	}
}
