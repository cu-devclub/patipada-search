package communication_test

import (
	"data-management/communication"
	"data-management/messages"
	mock "data-management/mock/communication"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockGRPC() communication.GRPCInterface {
	mockGrpc := mock.NewMockgRPC()
	return mockGrpc
}

func TestAuth(t *testing.T) {
	mockGrpc := setupMockGRPC()

	t.Run("Success Authorization", func(t *testing.T) {
		mock.SetAuthorizationResponse(true)
		token := "token"
		requiredRole := "role"
		res, err := mockGrpc.Authorization(token, requiredRole)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("Forbidden Authorization", func(t *testing.T) {
		mock.SetAuthorizationResponse(false)
		errMessage := "User is not authorized"
		token := "invalid_token"
		requiredRole := "role"
		res, err := mockGrpc.Authorization(token, requiredRole)

		assert.Equal(t, err.Error(), errMessage)
		assert.Equal(t, false, res)
	})
}

func TestSearch(t *testing.T) {
	mockGrpc := setupMockGRPC()

	t.Run("Success Search", func(t *testing.T) {
		mock.SetSearchResponse(true)
		recordID := "record_id"
		res, err := mockGrpc.SearchRecord(recordID)

		assert.Nil(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		mock.SetSearchResponse(false)
		errMessage := messages.RECORD_NOT_FOUND
		recordID := "invalid_record_id"
		res, err := mockGrpc.SearchRecord(recordID)

		assert.Equal(t, err.Error(), errMessage)
		assert.Equal(t, false, res)
	})
}
