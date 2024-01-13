package usecases

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/helper"
	"data-management/request/models"
	"log"
)

var statusArr = []string{"pending", "approved", "rejected"}

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
//   - *errors.RequestError: An error that occurred during the operation, if any.
//       - status of 400 or 500
func (r *requestUsecase) GetRequest(status, username, requestID, index, approvedBy string) ([]*models.Request, *errors.RequestError) {
	log.Println("Get Request with status", status, "username", username, "requestID", requestID, "index", index, "approvedBy", approvedBy)
	// validate status
	if status != "" && !helper.Contains(status, statusArr) {
		log.Println("Error validate status", status)
		return nil, errors.CreateError(400, messages.BAD_REQUEST)
	}
	// validate username
	if username != "" {
		result, err := r.requestRepositories.ValidateUsername(username)
		if err != nil || result == false {
			log.Println("Error validate username", username)
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// validate record ID
	if index != "" {
		result, err := r.requestRepositories.ValidateRecordIndex(username)
		if err != nil || result == false {
			log.Println("Error validate record index", index)
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// validate approved by
	if approvedBy != "" {
		result, err := r.requestRepositories.ValidateUsername(approvedBy)
		if err != nil || result == false {
			log.Println("Error validate approved by", approvedBy)
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// create filter
	filter := entities.Filter{
		Status:     status,
		By:         username,
		RequestID:  requestID,
		Index:      index,
		ApprovedBy: approvedBy,
	}

	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		log.Println("Error convert filter to bsonM", err)
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	log.Println("Get Request with filter", bsonFilter)
	// get request from repository
	entitiesRequests, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		log.Println("Error get request from repository", err)
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	var modelsRequests []*models.Request
	for _, entitiesRequest := range entitiesRequests {
		modelsRequest := helper.EntityToModels(entitiesRequest)
		modelsRequests = append(modelsRequests, modelsRequest)
	}

	log.Println("Get Request success with result", modelsRequests)
	return modelsRequests, nil
}
