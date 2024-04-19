package repositories

import "search-esdb-service/record/entities"

type MLRepository interface {
	Text2VecGateway(text string) ([]*entities.Text2VecResponse, error)
}
