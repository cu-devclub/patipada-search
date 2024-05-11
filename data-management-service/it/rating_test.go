package tests

import (
	"data-management/config"
	"data-management/logging"
	test_container_database "data-management/mock/testcontainer/database"
	ratingHandler "data-management/ratings/handlers"
	"data-management/ratings/models"
	ratingRepository "data-management/ratings/repositories"
	ratingUsecase "data-management/ratings/usecases"
	"data-management/testutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRatingTestEnvironment() (ratingHandler.Handlers, func()) {
	logging.NewSLogger()

	err := config.LoadConfig("../.")
	if err != nil {
		panic(err)
	}

	db, dbCleanUp, err := test_container_database.NewMockMongoClient()
	if err != nil {
		panic(err)
	}

	// rating Construct
	ratingRepositories := ratingRepository.NewRatingRepository(db)

	ratingUsecase := ratingUsecase.NewRatingUsecase(&ratingRepositories)

	ratingHandlers := ratingHandler.NewRatingHandler(&ratingUsecase)

	return ratingHandlers, dbCleanUp
}

func setupRatingGinEngine(handler ratingHandler.Handlers) *gin.Engine {
	g := gin.Default()
	g.POST("/ratings", handler.InsertRating)
	g.GET("/ratings", handler.GetRatings)
	g.GET("/ratings/average", handler.GetSummaryRatings)

	return g
}

func TestRating(t *testing.T) {
	handler, cleanup := setUpRatingTestEnvironment()
	defer cleanup()

	g := setupRatingGinEngine(handler)

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

	t.Run("Success Get Ratings : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/ratings", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Success Get Average Ratings : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/ratings/average", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
