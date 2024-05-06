package usecases

import (
	"data-management/constant"
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/helper"
	"data-management/request/models"
	"time"
)

func (r *requestUsecase) InsertRequest(request *models.Request) error {
	if err := r.validator.Validate(request); err != nil {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// validate record index
	result, err := r.requestRepositories.ValidateRecordIndex(request.Index)
	if err != nil || !result {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	result, err = r.requestRepositories.ValidateUsername(request.By)
	if err != nil || !result {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// populate request entity
	nextSeq, err := r.requestRepositories.GetNextRequestCounter()
	if err != nil {
		return  err
	}
	requestID := helper.GenerateRequestID(nextSeq)
	
	requestEntity := &entities.Request{
		RequestID:  requestID,
		Index:      request.Index,
		YoutubeURL: request.YoutubeURL,
		Question:   request.Question,
		Answer:     request.Answer,
		StartTime:  request.StartTime,
		EndTime:    request.EndTime,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     constant.REQUEST_STATUS_PENDING,
		By:         request.By,
		ApprovedBy: "",
	}

	// insert request entity
	objId, err := r.requestRepositories.InsertRequest(requestEntity)
	if err != nil {
		return err
	}

	// populate request model with request entity to return
	request.ID = objId
	request.RequestID = requestEntity.RequestID
	request.CreatedAt = requestEntity.CreatedAt
	request.UpdatedAt = requestEntity.UpdatedAt
	request.Status = requestEntity.Status

	return nil
}
