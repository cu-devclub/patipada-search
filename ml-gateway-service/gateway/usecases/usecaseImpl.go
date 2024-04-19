package usecases

import (
	"log"
	"ml-gateway-service/config"
	"ml-gateway-service/gateway/helper"
	"ml-gateway-service/gateway/models"
	"ml-gateway-service/gateway/repositories"
)

type UsecaseImpl struct {
	repository repositories.Repository
	mlConfig   *config.MlConfig
}

func NewUsecase(repositroy repositories.Repository, mlConfig *config.MlConfig) Usecase {
	return &UsecaseImpl{
		repository: repositroy,
		mlConfig:   mlConfig,
	}
}

func (u *UsecaseImpl) Text2Vec(text string) ([]*models.Text2VecResponse, error) {
	text2VecResponses := []*models.Text2VecResponse{}
	for _, api := range u.mlConfig.APIs {
		log.Println("Making request to API: ", api.Name)
		text2VecResponse, err := u.repository.MakingText2VecRequest(api, text)
		if err != nil {
			return nil, err
		}
		log.Println("Response from API: ", api.Name, " is: ", text2VecResponse)
		respModel := helper.ConvertText2VecResponseEntityToModel(text2VecResponse)
		respModel.ScoreWeight = api.ScoreWeight
		text2VecResponses = append(text2VecResponses, respModel)
	}
	return text2VecResponses, nil
}
