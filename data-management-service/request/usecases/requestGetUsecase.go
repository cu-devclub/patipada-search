package usecases

import (
	"data-management/constant"
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/helper"
	"data-management/request/models"
	"data-management/util"
)

func (r *requestUsecase) GetRequest(status, username, requestID, index, approvedBy string) ([]*models.Request, error) {
	var statusArr = []string{constant.REQUEST_STATUS_PENDING, constant.REQUEST_STATUS_REVIEWED}

	// validate status
	if status != "" && !util.Contains(status, statusArr) {
		return nil, errors.CreateError(400, messages.BAD_REQUEST)
	}
	// validate username
	if username != "" {
		result, err := r.requestRepositories.ValidateUsername(username)
		if err != nil || !result {
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// validate record ID
	if index != "" {
		result, err := r.requestRepositories.ValidateRecordIndex(index)
		if err != nil || !result {
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// validate approved by
	if approvedBy != "" {
		result, err := r.requestRepositories.ValidateUsername(approvedBy)
		if err != nil || !result {
			return nil, errors.CreateError(400, messages.BAD_REQUEST)
		}
	}

	// create filter
	filter := &entities.Filter{
		Status:     status,
		By:         username,
		RequestID:  requestID,
		Index:      index,
		ApprovedBy: approvedBy,
	}

	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		return nil, err
	}

	// get request from repository
	entitiesRequests, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		return nil, err
	}

	if len(entitiesRequests) == 0 {
		return []*models.Request{}, nil
	}

	var modelsRequests []*models.Request
	for _, entitiesRequest := range entitiesRequests {
		modelsRequest := helper.EntityToModels(entitiesRequest)
		modelsRequests = append(modelsRequests, modelsRequest)
	}

	return modelsRequests, nil
}

func (r *requestUsecase) GetLastestRequestOfRecord(index string) (*models.Request, error) {
	result, err := r.requestRepositories.ValidateRecordIndex(index)
	if err != nil || !result {
		return nil, errors.CreateError(400, messages.BAD_REQUEST)
	}

	// create filter
	filter := &entities.Filter{
		Index: index,
	}
	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	// get request from repository
	entitiesRequest, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		return nil, err
	}

	if len(entitiesRequest) == 0 {
		return &models.Request{}, nil
	}

	// get the lastest request from `updated_at` field
	var lastestRequest *entities.Request
	for _, request := range entitiesRequest {
		if lastestRequest == nil {
			lastestRequest = request
			continue
		}

		if request.UpdatedAt.After(lastestRequest.UpdatedAt) {
			lastestRequest = request
		}
	}

	modelsRequest := helper.EntityToModels(lastestRequest)
	return modelsRequest, nil
}
