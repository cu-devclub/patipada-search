package usecases

import (
	"data-management/errors"
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
	InsertRequest(request *models.Request) *errors.RequestError

	// GetAllRequests retrieves all request from the request repositories.
	//
	// Return
	// - if no error, return slice of model Request
	// - if no data return nil
	// - Error 500 if internal server error
	GetAllRequests() ([]*models.Request, *errors.RequestError)

	// GetRequestByRequestID retrieves a Request from the Request repositories by its requestID.
	// It takes a requestID as an argument.
	//
	// Return
	// - if no error, return model Request
	// - Error 404 if no Request found
	// - Error 500 if internal server error
	GetRequestByRequestID(requestID string) (*models.Request, *errors.RequestError)


	GetRequestByRecordIndex(index string) (*models.Request, *errors.RequestError)
	
	// UpdateRequest updates a request in the MongoDB collection.
	//	The function takes a pointer to a models.Request object as input. The Request object is first validated
	// and then converted to an entity using the helper.ModelsToEntity function. The UpdatedAt field of the entity
	// is set to the current time.
	//
	// Return
	// - Success : no error (nil)
	// - 400 : Bad request (validation error)
	// - 500 : internal server error
	UpdateRequest(request *models.Request) *errors.RequestError
}
