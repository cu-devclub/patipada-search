package mock_ratings

import (
	"data-management/ratings/entities"
	"data-management/ratings/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

var ratings []*entities.Rating

func init() {
	ratings = make([]*entities.Rating, 0)
}

func SetMockRatingsValue(rating []*entities.Rating) {
	ratings = rating
}

type mockRatingsRepository struct {
	mongo             *mongo.Client
	ratingsCollection *mongo.Collection
}

func NewMockRatingRepository() repositories.Repository {
	return &mockRatingsRepository{
		mongo:             nil,
		ratingsCollection: nil,
	}
}

func (r *mockRatingsRepository) InsertRating(rating *entities.Rating) (string, error) {
	objectIDString := "1234567890"

	return objectIDString, nil
}

func (r *mockRatingsRepository) GetRatings() ([]*entities.Rating, error) {
	return ratings, nil
}

func GetMockRatingConstant() *entities.Rating {
	return &entities.Rating{
		ID:       "1234567890",
		RatingID: "1",
		Stars:    5,
		Comment:  "Good",
	}
}