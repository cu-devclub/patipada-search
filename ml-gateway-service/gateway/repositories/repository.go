package repositories

import (
	"ml-gateway-service/config"
	"ml-gateway-service/gateway/entities"
)

type Repository interface {
	MakingText2VecRequest(externalAPI *config.ExternalAPI, text string) (*entities.Text2VecResponse, error)
}
