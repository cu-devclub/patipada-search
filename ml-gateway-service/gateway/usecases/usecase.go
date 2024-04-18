package usecases

import "ml-gateway-service/gateway/models"

type Usecase interface {
	Text2Vec(text string) ([]*models.Text2VecResponse, error)
}
