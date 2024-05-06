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

func (r *Rating) MockRating() {
	mockRatingID := "1234567890"
	r.RatingID = &mockRatingID
	r.Stars = 5
	r.Comment = "Good"
}
