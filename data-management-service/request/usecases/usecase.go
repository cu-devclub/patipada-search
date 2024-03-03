package usecases

import (
	"data-management/request/models"
)

type UseCase interface {

	// InsertRequest inserts a new request into the database.
	// It takes a pointer to a models.Request as an argument.
	// The function will return an error if the insertion fails, otherwise it will return nil.
	//
	// The function first validates the Request index using the ValidateRequestIndex method of the requestRepositories.
	// If the validation fails, it returns an error.
	//
	// Then, it generates a unique requestID for the new request using the helper.GenerateRequestID function.
	// If the generation of the requestID fails, it returns an internal server error.
	//
	// The function then populates a entities.Request struct with the data from the models.Request struct,
	// the generated requestID, and the current time for the CreatedAt and UpdatedAt fields.
	// The Status field is set to "pending".
	//
	// Error Status Codes:
	//     400: ERR_REQUEST_INDEX_NOT_EXISTS
	//     500: INTERNAL_SERVER_ERROR
	//
	// Then, it inserts the entities.Request struct into the database using the InsertRequest method of the requestRepositories.
	// If the insertion fails, it returns an error.
	//
	// Finally, it populates the original models.Request struct with the requestID, CreatedAt, UpdatedAt, and Status fields from the entities.Request struct.
	InsertRequest(request *models.Request) error

	// UpdateRequest updates a request in the MongoDB collection.
	//	The function takes a pointer to a models.Request object as input. The Request object is first validated
	// and then converted to an entity using the helper.ModelsToEntity function. The UpdatedAt field of the entity
	// is set to the current time.
	//
	// Return
	// - Success : no error (nil)
	// - 400 : Bad request (validation error)
	// - 500 : internal server error
	UpdateRequest(request *models.Request) error

	// GetRequest retrieves requests based on the provided parameters.
	// It validates the parameters, creates a filter from them, and then retrieves the requests from the repository.
	// If a parameter is an empty string, it will not be included in the filter.
	// If an error occurs during the operation, it will be returned along with a nil slice.
	//
	// Parameters:
	//
	//	status: The status of the requests to retrieve.
	//	username: The username associated with the requests to retrieve.
	//	requestID: The ID of the request to retrieve.
	//	index: The index of the request to retrieve.
	//	approvedBy: The username of the user who approved the requests to retrieve.
	//
	// Returns:
	//   - []*models.Request: A slice of pointers to the matching requests. If no requests match the filter, the slice will be empty.
	//   - error: An error that occurred during the operation, if any.
	//       - status of 400 or 500
	GetRequest(status, username, requestID, index, approvedBy string) ([]*models.Request, error)

	// GetLastestRequestOfRecord retrieves the latest request of a record based on the provided index.
	// It validates the index, creates a filter from it, and then retrieves the requests from the repository.
	// The function then finds the latest request based on the `updated_at` field.
	// If an error occurs during the operation, the function returns an error along with a nil pointer.
	// If no requests match the filter, the function returns a nil pointer and a nil error.
	//
	// Parameters:
	//   index: The index of the record.
	//
	// Returns:
	//   - *models.Request: A pointer to the latest request. If no requests match the filter, the pointer will be nil.
	//   - error: An error that occurred during the operation, if any.
	//          Possible status codes are
	//			400 (Bad Request) and
	//          500 (Internal Server Error).
	GetLastestRequestOfRecord(index string) (*models.Request, error)
}
