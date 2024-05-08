package usecases

import (
	"data-management/ratings/entities"
	"data-management/ratings/models"
	"data-management/ratings/repositories"
	"data-management/util"
	"math"
)

type UsecaseImpl struct {
	ratingsRepository repositories.Repository
}

func NewRatingUsecase(ratingsRepository *repositories.Repository) Usecase {
	return &UsecaseImpl{
		ratingsRepository: *ratingsRepository,
	}
}

func (u *UsecaseImpl) InsertRating(rating *models.Rating) (*models.Rating, error) {
	ratingID := util.GenerateUUID()

	ratingEntity := &entities.Rating{
		RatingID: ratingID,
		Stars:    rating.Stars,
		Comment:  rating.Comment,
	}

	u.ratingsRepository.InsertRating(ratingEntity)

	ratingModels := &models.Rating{
		RatingID: &ratingID,
		Stars:    ratingEntity.Stars,
		Comment:  ratingEntity.Comment,
	}

	return ratingModels, nil
}

func (u *UsecaseImpl) GetRatings() ([]*models.Rating, error) {
	ratings, err := u.ratingsRepository.GetRatings()
	if err != nil {
		return nil, err
	}

	var ratingModels []*models.Rating
	for _, rating := range ratings {
		ratingModel := &models.Rating{
			RatingID: &rating.RatingID,
			Stars:    rating.Stars,
			Comment:  rating.Comment,
		}

		ratingModels = append(ratingModels, ratingModel)
	}

	return ratingModels, nil
}

func (u *UsecaseImpl) GetSummaryRatings() (*models.SummaryRating, error) {
	ratings, err := u.ratingsRepository.GetRatings()
	if err != nil {
		return &models.SummaryRating{}, err
	}

	totalStars := 0
	for _, rating := range ratings {
		totalStars += rating.Stars
	}

	averageStars := float64(totalStars) / float64(len(ratings))
	averageStarsTwoDecimal := math.Round(averageStars*100) / 100
	percentage := averageStars * 20
	percentageTwoDecimal := math.Round(percentage*100) / 100

	summary := &models.SummaryRating{
		AverageStars: averageStarsTwoDecimal,
		TotalRatings: len(ratings),
		Percentage:   percentageTwoDecimal,
	}

	return summary, nil
}
