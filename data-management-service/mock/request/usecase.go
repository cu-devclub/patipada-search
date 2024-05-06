package mock_request

import (
	"data-management/request/models"
	"data-management/request/repositories"
	"data-management/request/usecases"
	validator "data-management/structValidator"
)

type mockUsecase struct {
	requestRepositories repositories.Repositories
	validator           validator.Validator
}

func NewMockUsecase() usecases.UseCase {
	return &mockUsecase{
		requestRepositories: nil,
		validator:           nil,
	}
}

func (r *mockUsecase) GetRequest(status, username, requestID, index, approvedBy string) ([]*models.Request, error) {

	var modelsRequests []*models.Request
	// mock
	modelsRequest := &models.Request{}
	modelsRequest.MockData()
	modelsRequests = append(modelsRequests, modelsRequest)

	return modelsRequests, nil
}

func (r *mockUsecase) GetLastestRequestOfRecord(index string) (*models.Request, error) {
	// mock
	modelsRequest := &models.Request{}
	modelsRequest.MockData()
	return modelsRequest, nil
}

func (r *mockUsecase) InsertRequest(request *models.Request) error {
	return nil
}

func (r *mockUsecase) SummaryData() (*models.Summary, error) {
	requestSummary := &models.RequestSummary{
		RequestAmount:  2,
		ReviewedAmount: 1,
		PendingAmount:  1,
	}

	recordSummary := &models.RecordSummary{
		RecordAmount:      10,
		YouTubeClipAmount: 8,
	}

	summary := &models.Summary{
		RequestSummary: requestSummary,
		RecordSummary:  recordSummary,
	}

	return summary, nil
}

func (r *mockUsecase) UpdateRequest(request *models.Request) error {
	return nil
}

func (r *mockUsecase) SyncRequestRecord(request *models.SyncRequestRecord) error {
	return nil
}

func (r *mockUsecase) SyncAllRequestRecords() error {
	return nil
}
