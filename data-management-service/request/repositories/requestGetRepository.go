package repositories

import (
	"context"
	"data-management/request/entities"

	"gopkg.in/mgo.v2/bson"
)

// GetRequest retrieves requests from the database based on the provided filter.
// The filter is a map where the key is the field name and the value is the value to match.
// If the filter is empty, all requests will be returned.
// If an error occurs during the operation, it will be returned along with a nil slice.
//
// Parameters:
//   filter: A map representing the filter to apply to the requests. The key is the field name and the value is the value to match.
//
// Returns:
//   []*entities.Request: A slice of pointers to the matching requests. If no requests match the filter, the slice will be empty.
//   error: An error that occurred during the operation, if any.
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

