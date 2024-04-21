package repositories

import (
	"search-esdb-service/communication"
	"search-esdb-service/record/entities"
)

type MLRepository interface {
	Text2VecGateway(text string) ([]*entities.Text2VecResponse, error)
}

type MLServiceRepository struct {
	comm *communication.Communication
}

func NewMLServiceRepository(comm *communication.Communication) MLRepository {
	return &MLServiceRepository{
		comm: comm,
	}
}
