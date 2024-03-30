package repositories

import "data-management/ratings/entities"

type Repository interface {
	InsertRating(rating *entities.Rating) (string, error)
	GetRatings() ([]*entities.Rating, error)
}
