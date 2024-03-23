package usecases

import "data-management/ratings/models"

type Usecase interface {
	InsertRating(rating *models.Rating) (*models.Rating, error)
	GetRatings() ([]*models.Rating, error)
	GetAverageRatings() (*models.SummaryRating, error)
}
