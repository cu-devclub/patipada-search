package repositories

import (
	"context"
	"data-management/request/entities"

	"gopkg.in/mgo.v2/bson"
)

// GetAllRequests retrieves all Requests from the MongoDB collection.
// It returns a slice of pointers to the requests ; if no data return nil
// and any error encountered.
//
// The function works as follows:
//  1. It calls the Find method on the RequestCollection with an empty filter (bson.M{}), which matches all documents in the collection.
//  2. If an error occurs during the Find operation, it returns the error and a nil slice.
//  3. It iterates over the cursor returned by the Find method. For each document in the cursor:
//     a. It decodes the document into a Request struct.
//     b. If an error occurs during decoding, it returns the error and a nil slice.
//     c. It appends the Request struct to the Requests slice.
//  4. After all documents have been processed, it returns the Requests slice and a nil error.
func (r *requestRepositories) GetAllRequests() ([]*entities.Request, error) {
	var requests []*entities.Request

	cursor, err := r.requestCollection.Find(context.Background(), bson.M{})
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

// GetRequestByRequestID retrieves a request from the request repositories by its request ID.
// It returns a pointer to the request and an error.
//
// The function works as follows:
// 1. It creates a filter that matches documents where the request_id field is equal to the provided requestID.
// 2. It calls the FindOne method on the requestCollection with the filter, which returns the first document that matches the filter.
// 3. It decodes the returned document into a request struct.
// 4. If an error occurs during the FindOne operation or the decoding, it returns a nil request and the error.
// 5. If no error occurs, it returns the Request and a nil error.
//
// If no document matches the filter, the function returns a nil Request and an error.
func (r *requestRepositories) GetRequestByRequestID(requestID string) (*entities.Request, error) {
	var request *entities.Request

	filter := bson.M{"request_id": requestID}
	err := r.requestCollection.FindOne(context.Background(), filter).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
