package models

type (
	Text2VecResponse struct {
		Name        string
		Embedding   []float32
		ScoreWeight float32
	}
)
