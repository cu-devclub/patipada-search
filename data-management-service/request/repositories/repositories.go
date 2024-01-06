package repositories

import (
	"data-management/request/entities"
)

type Repositories interface {

	// InsertRequest inserts a new Request into the MongoDB Request collection.
	// It takes a pointer to an entities.Request as an argument.
	// The function will return an objectID and error if the insertion fails, 
	// otherwise it will return "", nil.
	//
	// Usage:
	//
	//	Request := &entities.Request{...}
	//	err := RequestRepositories.InsertRequest(Request)
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	InsertRequest(request *entities.Request) (string, error)

	ValidateRecordIndex(request string) error

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
	//     request := &entities.Request{
	//         ID: "60d5ecf7c88f9a200f9e2c5a",
	//         Question: "Updated question",
	//         Answer: "Updated answer"
	//         ....
	//     }
	//     err := r.UpdateRequest(request)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	UpdateRequest(request *entities.Request) error

	// IncrementRequestCounter increments the sequence number in the "counters" collection in MongoDB.
	// This function is used to generate a unique sequence number for each new Request.
	// It returns the next sequence number and any error encountered.
	//
	// The function works by finding the document in the "counters" collection with _id "Request",
	// incrementing its "seq" field, and returning the updated "seq" value.
	// If the document does not exist, MongoDB will return an error.
	//
	// Usage:
	//     nextSeq, err := RequestRepositories.IncrementRequestCounter()
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     nextRequestID := fmt.Sprintf("REQ%d", nextSeq)
	//
	// This will generate a unique RequestID for each new Request, like "REQ1", "REQ2", etc.
	IncrementRequestCounter() (int, error)

	// GetAllRequests retrieves all Requests from the MongoDB collection.
	// It returns a slice of pointers to the Requests ; if no data return nil
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
	GetAllRequests() ([]*entities.Request, error)

	// GetRequestByRequestID retrieves a Request from the Request repositories by its request ID.
	// It returns a pointer to the Request and an error.
	//
	// The function works as follows:
	// 1. It creates a filter that matches documents where the request_id field is equal to the provided requestID.
	// 2. It calls the FindOne method on the RequestCollection with the filter, which returns the first document that matches the filter.
	// 3. It decodes the returned document into a Request struct.
	// 4. If an error occurs during the FindOne operation or the decoding, it returns a nil Request and the error.
	// 5. If no error occurs, it returns the Request and a nil error.
	//
	// If no document matches the filter, the function returns a nil Request and an error.
	GetRequestByRequestID(requestID string) (*entities.Request, error)
}
