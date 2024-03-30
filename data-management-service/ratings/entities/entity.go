package entities

type (
	Rating struct {
		ID       string `bson:"_id,omitempty" json:"id,omitempty"`
		RatingID string `bson:"rating_id"`
		Stars    int    `bson:"stars"`
		Comment  string `bson:"comment"`
	}
)
