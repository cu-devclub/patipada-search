package models

type (
	Rating struct {
		RatingID *string `json:"rating_id"`
		Stars    int     `json:"stars"`
		Comment  string  `json:"comment"`
	}

	SummaryRating struct {
		AverageStars float64 `json:"average_stars"`
		TotalRatings int     `json:"total_ratings"`
		Percentage   float64 `json:"percentage_rating"`
	}
)
