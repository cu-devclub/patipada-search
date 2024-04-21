package entities

type (
	Text2VecResponse struct {
		Name      string    `json:"name"`
		Embedding []float32 `json:"embedding"`
	}
)
