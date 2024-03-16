package repositories

type MLServiceRepository struct {}

func NewMLServiceRepository() MLRepository {
	return &MLServiceRepository{}
}