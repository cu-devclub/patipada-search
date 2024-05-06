package request_repo_test

import (
	mock_communication "data-management/mock/communication"
	test_container_database "data-management/mock/testcontainer/database"
	"data-management/request/entities"
	"data-management/request/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMongoRepoTest() (repositories.Repositories, func()) {
	// setup
	db, cleanup, err := test_container_database.NewMockMongoClient()
	if err != nil {
		panic(err)
	}

	comm := mock_communication.MockCommunication()

	repo := repositories.NewRepositories(db, &comm)
	return repo, cleanup
}

func TestMongoRepo(t *testing.T) {
	repo, cleanup := setupMongoRepoTest()
	defer cleanup()

	t.Run("Get Request no value", func(t *testing.T) {
		filter := entities.Filter{}
		filterBson, err := filter.ConvertToBsonM()
		assert.Nil(t, err)

		requests, err := repo.GetRequest(filterBson)
		assert.Nil(t, err)

		assert.Equal(t, 0, len(requests))
	})

	t.Run("InsertRequest", func(t *testing.T) {
		request := &entities.Request{}
		request.MockData()

		res, err := repo.InsertRequest(request)
		assert.Nil(t, err)

		assert.NotEqual(t, "", res)
	})

	t.Run("GetRequest the value exist", func(t *testing.T) {
		expectedRequest := &entities.Request{}
		expectedRequest.MockData()
		expectedRequest.RequestID = "mock"

		// insert the request
		_, err := repo.InsertRequest(expectedRequest)
		assert.Nil(t, err)

		// get the request
		filter := entities.Filter{
			RequestID: expectedRequest.RequestID,
		}
		filterBson, err := filter.ConvertToBsonM()
		assert.Nil(t, err)

		requests, err := repo.GetRequest(filterBson)
		assert.Nil(t, err)

		// the requests can have multiple we want to check that the request we inserted is in the list
		found := false
		for _, request := range requests {
			if request.RequestID == expectedRequest.RequestID {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("Get Request with filter", func(t *testing.T) {
		expectedRequest := &entities.Request{}
		expectedRequest.MockData()
		expectedRequest.RequestID = "mock-test-data"

		// insert the request
		_, err := repo.InsertRequest(expectedRequest)
		assert.Nil(t, err)

		// get the request
		filter := entities.Filter{
			RequestID: expectedRequest.RequestID,
		}
		filterBson, err := filter.ConvertToBsonM()
		assert.Nil(t, err)

		requests, err := repo.GetRequest(filterBson)
		assert.Nil(t, err)

		assert.Equal(t, 1, len(requests))
		assert.Equal(t, expectedRequest.RequestID, requests[0].RequestID)
	})

	t.Run("Update Request", func(t *testing.T) {
		expectedRequest := &entities.Request{}
		expectedRequest.MockData()
		expectedRequest.RequestID = "mock-test-data-2"

		// insert the request
		objectID, err := repo.InsertRequest(expectedRequest)
		assert.Nil(t, err)

		// update the request
		expectedRequest.ID = objectID
		expectedRequest.Status = "reviewed"
		err = repo.UpdateRequest(expectedRequest)
		assert.Nil(t, err)

		// get the request
		filter := entities.Filter{
			RequestID: expectedRequest.RequestID,
		}
		filterBson, err := filter.ConvertToBsonM()
		assert.Nil(t, err)

		requests, err := repo.GetRequest(filterBson)
		assert.Nil(t, err)

		assert.Equal(t, 1, len(requests))
		assert.Equal(t, expectedRequest.Status, requests[0].Status)
	})

	t.Run("Get Next Request Counter", func(t *testing.T) {
		counter, err := repo.GetNextRequestCounter()
		assert.Nil(t, err)

		assert.Equal(t, 1, counter)
	})

	t.Run("Get Record counter", func(t *testing.T) {
		_, err := repo.GetRecordCounter()
		assert.Nil(t, err)
	})

	t.Run("Upsert Record Counter", func(t *testing.T) {
		counter := &entities.RecordCounter{
			RecordAmount:      12,
			YoutubeClipAmount: 10,
		}
		err := repo.UpsertRecordCounter(counter)
		assert.Nil(t, err)
	})

	t.Run("Upsert Record Counter and Get Record Counter", func(t *testing.T) {
		recordCounter := &entities.RecordCounter{
			RecordAmount:      13,
			YoutubeClipAmount: 19,
		}
		err := repo.UpsertRecordCounter(recordCounter)
		assert.Nil(t, err)

		counter, err := repo.GetRecordCounter()
		assert.Nil(t, err)

		assert.Equal(t, 13, counter.RecordAmount)
		assert.Equal(t, 19, counter.YoutubeClipAmount)
	})
}
