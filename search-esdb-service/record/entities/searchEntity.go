package entities

type (

	KeywordSearchStruct struct {
		Query                    string        `json:"Query"`
		KeywordSearchFields      []string      `json:"KeywordSearchFields"`
		Config                   *SearchConfig `json:"Config"`
	}

	VectorSearchStruct struct {
		IndexName    string              `json:"IndexName"`
		VectorFields []*Text2VecResponse `json:"VectorFields"`
		Config       *SearchConfig       `json:"Config"`
	}
	HybridSearchStruct struct {
		Query                    string              `json:"Query"`
		KeywordSearchFields      []string            `json:"KeywordSearchFields"`
		KeywordSearchScoreWeight float64             `json:"KeywordSearchScoreWeight"`
		VectorFields             []*Text2VecResponse `json:"VectorFields"`
		Config                   *SearchConfig       `json:"Config"`
	}

	SearchConfig struct {
		IndexName   string `json:"IndexName"`
		Offset      int    `json:"Offset"`
		Amount      int    `json:"Amount"`
		CountNeeded bool   `json:"CountNeeded"`
	}

	Text2VecResponse struct {
		Embedding   []float32 `json:"Embedding"`
		Name        string    `json:"Name"`
		ScoreWeight float32   `json:"ScoreWeight"`
	}
)
