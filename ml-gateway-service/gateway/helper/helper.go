package helper

import (
	"ml-gateway-service/gateway/entities"
	"ml-gateway-service/gateway/models"
)

func ConvertText2VecResponseEntityToModel(text2VecResponseEntity *entities.Text2VecResponse) *models.Text2VecResponse {
	return &models.Text2VecResponse{
		Name:      text2VecResponseEntity.Name,
		Embedding: text2VecResponseEntity.Embedding,
	}
}
