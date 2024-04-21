package repositories

import (
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
)

func (r *MLServiceRepository) Text2VecGateway(text string) ([]*entities.Text2VecResponse, error) {
	// call grpc method
	com := *r.comm
	grpcResponse, err := com.Text2Vec(text)
	if err != nil {
		return nil, err
	}

	text2VecResponse := helper.ConvertgRPCText2VecResonseToEntityResponses(grpcResponse)

	return text2VecResponse, nil
}
