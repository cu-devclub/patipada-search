package usecases

import (
	"data-management/constant"
	"data-management/errors"
	"data-management/messages"
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
	for _, req := range requests {
		if req.UpdatedAt.After(request.UpdatedAt) {
			continue
		}

		if req.ID == request.ID {
			continue
		}

		req.Status = constant.STATUS_REVIEWED
		requestEntitiy := helper.ModelsToEntity(req)
		requestEntitiy.UpdatedAt = time.Now()
		if err := r.requestRepositories.UpdateRequest(requestEntitiy); err != nil {
			return err
		}
	}

	// --- Update the current request
	requestEntity := helper.ModelsToEntity(request)
	requestEntity.UpdatedAt = time.Now()
	requestEntity.Status = constant.STATUS_REVIEWED
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
