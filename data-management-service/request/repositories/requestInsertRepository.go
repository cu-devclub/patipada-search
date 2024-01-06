package repositories

import (
	"data-management/errors"
	"data-management/request/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

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
