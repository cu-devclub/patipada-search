package mock_ratings

import (
	"data-management/ratings/models"
	"data-management/ratings/repositories"
	"data-management/ratings/usecases"
)

type MockUsecase struct {
	ratingsRepository repositories.Repository
}

func NewMockRatingUsecase() usecases.Usecase {
	return &MockUsecase{
		ratingsRepository: nil,
	}
}

func (u *MockUsecase) InsertRating(rating *models.Rating) (*models.Rating, error) {
	ratingID := "1234567890"
	ratingModels := &models.Rating{
		RatingID: &ratingID,
		Stars:    rating.Stars,
		Comment:  rating.Comment,
	}

	return ratingModels, nil
}

func (u *MockUsecase) GetRatings() ([]*models.Rating, error) {
	ratingModels := make([]*models.Rating,0)
	
	return ratingModels, nil
}

func (u *MockUsecase) GetAverageRatings() (*models.SummaryRating, error) {

	summary := &models.SummaryRating{
		AverageStars: 4.32,
		TotalRatings: 5,
	}

	return summary, nil
}
