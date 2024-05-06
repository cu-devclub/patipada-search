package usecases

import (
	"data-management/constant"
	"data-management/errors"
	"data-management/messages"
	"data-management/request/entities"
	"data-management/request/models"
)

func (r *requestUsecase) SummaryData() (*models.Summary, error) {
	filter := &entities.Filter{}
	bsonFilter, err := filter.ConvertToBsonM()
	if err != nil {
		return nil, errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}
	requests, err := r.requestRepositories.GetRequest(bsonFilter)
	if err != nil {
		return nil, err
	}

	totalRequest := len(requests)

	reviewedAmount := 0
	pendingAmount := 0

	for _, request := range requests {
		if request.Status == constant.REQUEST_STATUS_REVIEWED {
			reviewedAmount++
		} else if request.Status == constant.REQUEST_STATUS_PENDING {
			pendingAmount++
		}
	}

	requestSummary := &models.RequestSummary{
		RequestAmount:  totalRequest,
		ReviewedAmount: reviewedAmount,
		PendingAmount:  pendingAmount,
	}

	recordsCounter, err := r.requestRepositories.GetRecordCounter()
	if err != nil {
		return nil, err
	}

	if recordsCounter == nil {
		recordsCounter = &entities.RecordCounter{
			RecordAmount:      0,
			YoutubeClipAmount: 0,
		}
	}

	recordSummary := &models.RecordSummary{
		RecordAmount:      recordsCounter.RecordAmount,
		YouTubeClipAmount: recordsCounter.YoutubeClipAmount,
	}

	summary := &models.Summary{
		RequestSummary: requestSummary,
		RecordSummary:  recordSummary,
	}

	return summary, nil
}
