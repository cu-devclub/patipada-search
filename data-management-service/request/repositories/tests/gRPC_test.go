package request_repo_test

import (
	"data-management/communication"
	mock_communication "data-management/mock/communication"
	"data-management/request/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupRepoGRPCTest() repositories.Repositories {

	gRPC := mock_communication.NewMockgRPC()
	rabbitMQ := mock_communication.MockRabbitMQ()

	comm := communication.NewCommunicationImpl(gRPC, rabbitMQ)

	repo := repositories.NewRepositories(&mongo.Client{}, &comm)

	return repo
}

func TestValidateRecordIndex(t *testing.T) {
	repo := setupRepoGRPCTest()

	t.Run("RecordID is empty", func(t *testing.T) {
		result, err := repo.ValidateRecordIndex("")
		assert.False(t, result)
		assert.Nil(t, err)
	})

	t.Run("RecordID is not empty and valid", func(t *testing.T) {
		mock_communication.SetSearchResponse(true)
		result, err := repo.ValidateRecordIndex("123")
		assert.True(t, result)
		assert.Nil(t, err)
	})

	t.Run("RecordID is not empty and invalid", func(t *testing.T) {
		mock_communication.SetSearchResponse(false)
		result, err := repo.ValidateRecordIndex("123")
		assert.False(t, result)
		assert.NotNil(t, err)
	})
}

func TestValidateUsername(t *testing.T){
	repo := setupRepoGRPCTest()

	t.Run("Username is valid", func(t *testing.T){
		mock_communication.SetVerifyUsernameResponse(true)
		result, err := repo.ValidateUsername("valid")
		assert.True(t, result)
		assert.Nil(t, err)
	})

	t.Run("Username is invalid", func(t *testing.T){
		mock_communication.SetVerifyUsernameResponse(false)
		result, err := repo.ValidateUsername("invalid")
		assert.False(t, result)
		assert.NotNil(t, err)
	})
}