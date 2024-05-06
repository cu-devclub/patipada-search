package rating_hadler_tests

import (
	mock_ratings "data-management/mock/ratings"
	"data-management/ratings/handlers"
	"data-management/ratings/models"
	"data-management/testutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRatingHandler() *gin.Engine {
	ratingUsecase := mock_ratings.NewMockRatingUsecase()
	handler := handlers.NewRatingHandler(&ratingUsecase)

	g := gin.Default()
	g.POST("/ratings", handler.InsertRating)
	g.GET("/ratings", handler.GetRatings)
	g.GET("/ratings/average", handler.GetAverageRatings)

	return g
}

func TestInsertRequest(t *testing.T) {
	g := setUpRatingHandler()
	
	t.Run("Success Insert Rating : 201 Created", func(t *testing.T) {
		ratingModel := models.Rating{}
		ratingModel.MockRating()

		req := testutil.CreateNewRequest("POST", "/ratings", ratingModel)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("Fail Insert Rating : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/ratings", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestGetRatings(t *testing.T) {
	g := setUpRatingHandler()
	
	t.Run("Success Get Ratings : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/ratings", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

func TestGetAverageRatings(t *testing.T) {
	g := setUpRatingHandler()
	
	t.Run("Success Get Average Ratings : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/ratings/average", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
