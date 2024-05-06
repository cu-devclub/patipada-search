package mock_request

import (
	"data-management/communication"
	"data-management/request/entities"
	"data-management/request/repositories"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var requests []*entities.Request

var validateRecordResponse bool

var recordCounterResponse *entities.RecordCounter

func init() {
	requests = make([]*entities.Request, 0)
}

func SetMockRequestsValue(request []*entities.Request) {
	requests = request
}

func SetValidateRecordResponse(response bool) {
	validateRecordResponse = response
}

func SetRecordCounterResponse(response *entities.RecordCounter) {
	recordCounterResponse = response
}

type mockRepository struct {
	mongo                    *mongo.Client
	requestCollection        *mongo.Collection
	requestCounterCollection *mongo.Collection
	communicationClient      communication.Communication
}

func NewMockRepositories() repositories.Repositories {
	return &mockRepository{
		mongo:                    nil,
		requestCollection:        nil,
		requestCounterCollection: nil,
		communicationClient:      nil,
	}
}

func (r *mockRepository) GetRequest(filter bson.M) ([]*entities.Request, error) {
	return requests, nil
}

func (r *mockRepository) InsertRequest(request *entities.Request) (string, error) {
	objectIDString := "60d5ecf7c88f9a200f9e2c5a"

	return objectIDString, nil
}

func (r *mockRepository) UpdateRequest(request *entities.Request) error {
	return nil
}

func (r *mockRepository) ValidateRecordIndex(recordID string) (bool, error) {
	return validateRecordResponse, nil
}

func (r *mockRepository) ValidateUsername(username string) (bool, error) {
	return true, nil
}

func (r *mockRepository) UpdateRecord(record *entities.Record) error {
	return nil
}

func (r *mockRepository) GetNextRequestCounter() (int, error) {
	return 1, nil
}

func (r *mockRepository) UpsertRecordCounter(recordCounter *entities.RecordCounter) error {
	return nil
}

func (r *mockRepository) GetRecordCounter() (*entities.RecordCounter, error) {
	return recordCounterResponse, nil
}

func GetMockRequestConstant() *entities.Request {
	return &entities.Request{
		ID:         "60d5ecf7c88f9a200f9e2c5a",
		RequestID:  "REQ1",
		Index:      "sadads-1",
		YoutubeURL: "youtubeURL",
		Question:   "question",
		Answer:     "answer",
		StartTime:  "startTime",
		EndTime:    "endTime",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "pending",
		By:         "username",
		ApprovedBy: "approvedBy",
	}
}
