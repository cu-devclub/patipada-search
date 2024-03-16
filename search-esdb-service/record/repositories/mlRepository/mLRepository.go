package repositories

type MLRepository interface {
	RemoveStopWordFromTokensArrays(tokens []string) ([]string, error)

	RemoveStopWordFromQuery(query string) ([]string, error)

	TokenizeQuery(query string) ([]string, error)

	PerformLDATopicModelling(tokens []string) ([]float64, error)
}
