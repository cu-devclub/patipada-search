package usecases

import (
	"data-management/constant"
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/helper"
	"data-management/request/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (r *requestUsecase) UpdateRequest(request *models.Request) error {
	// validate the request

	if err := r.validator.Validate(request); err != nil {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid record index
	result, err := r.requestRepositories.ValidateRecordIndex(request.Index)
	if err != nil || !result {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid by
	result, err = r.requestRepositories.ValidateUsername(request.By)
	if err != nil || !result {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// check valid approved by
	result, err = r.requestRepositories.ValidateUsername(request.ApprovedBy)
	if err != nil || !result {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// ---  Get all requests that has the same record Index by Get usecase
	requests, er := r.GetRequest("", "", "", request.Index, "")
	if er != nil {
		return er
	}

	// ---  Update all request that come before the current request by setting status to "reviewed"
	previousRequest := helper.UpdatePreviousRequestsStatus(requests, request)
	for _, req := range previousRequest {
		if err := r.requestRepositories.UpdateRequest(req); err != nil {
			return err
		}
	}

	// --- Update the current request
	requestEntity := helper.ModelsToEntity(request)
	requestEntity.UpdatedAt = time.Now()
	requestEntity.Status = constant.REQUEST_STATUS_REVIEWED
	if err := r.requestRepositories.UpdateRequest(requestEntity); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.CreateError(400, messages.BAD_REQUEST)
		}
		return err
	}

	recordEntity := helper.RequestToRecordsEntity(requestEntity)
	recordEntity.ExtractHTML()

	// --- Update the record by sending plain text of `startTime`,`endTime`,`question`,`answer` to the record (search) service
	err = r.requestRepositories.UpdateRecord(recordEntity)
	if err != nil {
		return err
	}

	return nil
}



func (r *requestUsecase) SyncRequestRecord(request *models.SyncRequestRecord) error {
	filter := &entities.Filter{
		RequestID: request.RequestId,
	}

	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		return err
	}

	requests, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		return err
	}

	if len(requests) == 0 {
		return errors.CreateError(404, messages.NOT_FOUND)
	}

	if len(requests) > 1 {
		return errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	requestFetched := requests[0]

	if requestFetched.Status != constant.REQUEST_STATUS_REVIEWED {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	recordEntity := helper.RequestToRecordsEntity(requestFetched)
	recordEntity.ExtractHTML()

	err = r.requestRepositories.UpdateRecord(recordEntity)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestUsecase) SyncAllRequestRecords() error {
	filter := &entities.Filter{
		Status: constant.REQUEST_STATUS_REVIEWED,
	}
	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		return err
	}

	requests, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		return err
	}

	// Get the latest one for each record index sort by request ID
	latestRequests := make(map[string]*entities.Request)
	for _, request := range requests {
		if _, ok := latestRequests[request.Index]; !ok {
			latestRequests[request.Index] = request
			continue
		}

		if latestRequests[request.Index].ID < request.ID {
			latestRequests[request.Index] = request
		}
	}

	for _, request := range latestRequests {
		recordEntity := helper.RequestToRecordsEntity(request)
		recordEntity.ExtractHTML()

		err = r.requestRepositories.UpdateRecord(recordEntity)
		if err != nil {
			return err
		}
	}

	return nil
}


