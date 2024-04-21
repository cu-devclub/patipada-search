package repositories

import (
	"encoding/json"
	"fmt"
	"ml-gateway-service/config"
	"ml-gateway-service/errors"
	"ml-gateway-service/gateway/entities"
	"ml-gateway-service/util"
	"net/url"
)

type gatewayRepository struct {
}

func NewGatewayRepository() Repository {
	return &gatewayRepository{}
}

func (r *gatewayRepository) MakingText2VecRequest(externalAPI *config.ExternalAPI, text string) (*entities.Text2VecResponse, error) {
	encodedText := url.QueryEscape(text)
	resp, err := util.HttpGETRequest(externalAPI.URL + "?text=" + encodedText)
	if err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error calling external serivce: %v", err))
	}
	var text2VecResponse *entities.Text2VecResponse
	err = json.Unmarshal(*resp, &text2VecResponse)
	if err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error unmarshalling response: %v with %v", resp, err))
	}

	return text2VecResponse, nil
}
