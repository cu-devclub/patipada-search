package repositories

import (
	"context"
	"data-management/communication"
	"data-management/errors"
	"data-management/request/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type requestRepositories struct {
	mongo                    *mongo.Client
	requestCollection        *mongo.Collection
	requestCounterCollection *mongo.Collection
	communicationClient      communication.Communication
}

func NewRequestRepositories(mongo *mongo.Client, c *communication.Communication) Repositories {
	requestCollection := mongo.Database("data").Collection("request")
	requestCounterCollection := mongo.Database("data").Collection("counters")
	return &requestRepositories{
		mongo:                    mongo,
		requestCollection:        requestCollection,
		requestCounterCollection: requestCounterCollection,
		communicationClient:      *c,
	}
}

// GetRequest retrieves requests from the database based on the provided filter.
// The filter is a map where the key is the field name and the value is the value to match.
// If the filter is empty, all requests will be returned.
// If an error occurs during the operation, it will be returned along with a nil slice.
//
// Parameters:
//
//	filter: A map representing the filter to apply to the requests. The key is the field name and the value is the value to match.
//
// Returns:
//
//	[]*entities.Request: A slice of pointers to the matching requests. If no requests match the filter, the slice will be empty.
//	error: An error that occurred during the operation, if any.
func (r *requestRepositories) GetRequest(filter bson.M) ([]*entities.Request, error) {
	var requests []*entities.Request

	cursor, err := r.requestCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var request *entities.Request
		err := cursor.Decode(&request)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}

// InsertRequest inserts a new Request into the MongoDB Request collection.
// It takes a pointer to an entities.Request as an argument.
// The function will return an objectID and error if the insertion fails,
// otherwise it will return "", nil.
// Usage:
//
//	Request := &entities.Request{...}
//	err := RequestRepositories.InsertRequest(Request)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *requestRepositories) InsertRequest(request *entities.Request) (string, error) {
	result, err := r.requestCollection.InsertOne(context.TODO(), request)
	if err != nil {
		return "", err
	}

	// Get the inserted ID and assert it to an ObjectID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.CreateError(500, "Cannot assert inserted ID to an ObjectID")
	}

	// Convert the ObjectID to a string
	objectIDString := objectID.Hex()

	return objectIDString, nil
}

// UpdateRequest updates a request in the MongoDB collection.
//
// The function takes a pointer to an entities.Request object as input. The ID field of the Request object
// is used to find the document to update in the MongoDB collection. The other fields of the Request object
// are used to update the corresponding fields in the document.
//
// The function returns an error if the update operation fails. If the update operation is successful,
// the function returns nil.
//
// The function uses the UpdateOne method from the mongo package to perform the update operation. The UpdateOne
// method updates the first document that matches the filter in the MongoDB collection.
//
// The function uses a context with a timeout of 10 seconds to ensure that the update operation is cancelled
// if it takes too long.
//
// Example:
//
//	request := &entities.Request{
//	    ID: "60d5ecf7c88f9a200f9e2c5a",
//	    Question: "Updated question",
//	    Answer: "Updated answer"
//	    ....
//	}
//	err := r.UpdateRequest(request)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *requestRepositories) UpdateRequest(request *entities.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"question":    request.Question,
			"answer":      request.Answer,
			"startTime":   request.StartTime,
			"endTime":     request.EndTime,
			"status":      request.Status,
			"updated_at":  request.UpdatedAt,
			"approved_by": request.ApprovedBy,
		},
	}

	_, err = r.requestCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
