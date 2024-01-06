package usecases

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/helper"
	"data-management/request/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllRequest retrieves all requests from the request repositories.
//
// Return
// - if no error, return slice of model request
// - if no data return nil
// - Error 500 if internal server error
func (r *requestUsecase) GetAllRequests() ([]*models.Request, *errors.RequestError) {
	entitiesRequests, err := r.requestRepositories.GetAllRequests()
	if err != nil {
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	var modelsRequests []*models.Request
	for _, entitiesRequest := range entitiesRequests {
		modelsRequest := helper.EntityToModels(entitiesRequest)
		modelsRequests = append(modelsRequests, modelsRequest)
	}

	return modelsRequests, nil
}

// GetRequestByRequestID retrieves a request from the request repositories by its RequestID.
// It takes a RequestID as an argument.
//
// Return
// - if no error, return model request
// - Error 404 if no request found
// - Error 500 if internal server error
func (r *requestUsecase) GetRequestByRequestID(requestID string) (*models.Request, *errors.RequestError) {
	entitiesRequest, err := r.requestRepositories.GetRequestByRequestID(requestID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.CreateError(404, messages.ERR_REQUEST_NOT_EXISTS)
		}
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	modelsRequest := helper.EntityToModels(entitiesRequest)

	return modelsRequest, nil
}
