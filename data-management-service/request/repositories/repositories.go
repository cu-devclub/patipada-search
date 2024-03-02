package repositories

import (
	"data-management/request/entities"

	"gopkg.in/mgo.v2/bson"
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

	ValidateRecordIndex(request string) (bool, error)

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
	GetRequest(filter bson.M) ([]*entities.Request, error)

	// ValidateUsername checks if the provided username is valid.
	// It uses the communication client's VerifyUsername method to perform the check.
	// If an error occurs during the operation, it will be returned along with a false boolean.
	//
	// Parameters:
	//   username: The username to validate.
	//
	// Returns:
	//   bool: A boolean indicating whether the username is valid. True if the username is valid, false otherwise.
	//   error: An error that occurred during the operation, if any.
	ValidateUsername(username string) (bool, error)

	UpdateRecord(record *entities.Record) error
}
