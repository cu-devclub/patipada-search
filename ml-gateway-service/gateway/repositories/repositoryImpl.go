package repositories

import (
	"ml-gateway-service/config"
	"ml-gateway-service/gateway/entities"
)

type gatewayRepository struct {
}

func NewGatewayRepository() Repository {
	return &gatewayRepository{}
}

func (r *gatewayRepository) MakingText2VecRequest(externalAPI *config.ExternalAPI, text string) (*entities.Text2VecResponse, error) {
	// resp, err := util.HttpGETRequest(externalAPI.URL + "?text=" + text)
	// if err != nil {
	// 	return nil, err
	// }

	// var text2VecResponse *entities.Text2VecResponse
	// err = json.Unmarshal(*resp, &text2VecResponse)
	// if err != nil {
	// 	return nil, err
	// }

	// ! Mock response
	// TODO : Remove this to call real API
	text2VecResponse := &entities.Text2VecResponse{
		Name:      "test",
		Embedding: []float32{1.0, 2.0, 3.0},
	}

	return text2VecResponse, nil
}
