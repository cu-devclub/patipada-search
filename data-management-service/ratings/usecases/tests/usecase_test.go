package rating_usecase_test

import (
	mock_ratings "data-management/mock/ratings"
	"data-management/ratings/entities"
	"data-management/ratings/models"
	"data-management/ratings/usecases"
	"data-management/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupUsecase() usecases.Usecase {
	mockRepo := mock_ratings.NewMockRatingRepository()

	// Create the usecase
	usecase := usecases.NewRatingUsecase(&mockRepo)

	return usecase
}
func TestInsertRating(t *testing.T) {
	usecase := setupUsecase()

	t.Run("InsertRating", func(t *testing.T) {
		rating := models.Rating{
			Stars:   5,
			Comment: "Good",
		}
		res, err := usecase.InsertRating(&rating)
		assert.Nil(t, err)

		assert.True(t, testutil.IsValidUUID(*res.RatingID))
	})
}

func TestGetRatings(t *testing.T) {
	usecase := setupUsecase()

	t.Run("GetRatings one value", func(t *testing.T) {

		expectedRating := mock_ratings.GetMockRatingConstant()

		mock_ratings.SetMockRatingsValue([]*entities.Rating{expectedRating})

		ratings, err := usecase.GetRatings()
		assert.Nil(t, err)

		assert.Equal(t, 1, len(ratings))
		assert.Equal(t, expectedRating.Stars, ratings[0].Stars)
		assert.Equal(t, expectedRating.Comment, ratings[0].Comment)
	})

	t.Run("GetRatings multiple values", func(t *testing.T) {
		expectedRating1 := mock_ratings.GetMockRatingConstant()
		expectedRating2 := mock_ratings.GetMockRatingConstant()
		expectedRating2.RatingID = "2"

		mock_ratings.SetMockRatingsValue([]*entities.Rating{expectedRating1, expectedRating2})

		ratings, err := usecase.GetRatings()
		assert.Nil(t, err)

		assert.Equal(t, 2, len(ratings))
		assert.Equal(t, expectedRating1.Stars, ratings[0].Stars)
		assert.Equal(t, expectedRating1.Comment, ratings[0].Comment)
		assert.Equal(t, expectedRating2.Stars, ratings[1].Stars)
		assert.Equal(t, expectedRating2.Comment, ratings[1].Comment)
	})
}

func TestGetAverageRating(t *testing.T) {
	usecase := setupUsecase()

	t.Run("GetAverageRating one value", func(t *testing.T) {
		expectedRating := mock_ratings.GetMockRatingConstant()

		mock_ratings.SetMockRatingsValue([]*entities.Rating{expectedRating})

		sumRating, err := usecase.GetAverageRatings()

		assert.Nil(t, err)
		assert.Equal(t, 1, sumRating.TotalRatings)
		assert.Equal(t, float64(expectedRating.Stars), sumRating.AverageStars)
	})

	t.Run("GetAverageRating multiple values", func(t *testing.T) {
		expectedRating1 := mock_ratings.GetMockRatingConstant()
		expectedRating2 := mock_ratings.GetMockRatingConstant()
		expectedRating2.RatingID = "2"

		mock_ratings.SetMockRatingsValue([]*entities.Rating{expectedRating1, expectedRating2})

		sumRating, err := usecase.GetAverageRatings()

		assert.Nil(t, err)
		assert.Equal(t, 2, sumRating.TotalRatings)
		assert.Equal(t, float64(expectedRating1.Stars+expectedRating2.Stars)/2, sumRating.AverageStars)
	})
}