package usecases

import "data-management/ratings/models"

type Usecase interface {
	InsertRating(rating *models.Rating) (*models.Rating, error)
	GetRatings() ([]*models.Rating, error)
	GetSummaryRatings() (*models.SummaryRating, error)
}
