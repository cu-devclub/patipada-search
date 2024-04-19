package communication

import (
	"context"
	"fmt"
	"net/http"
	"search-esdb-service/config"
	"search-esdb-service/errors"
	"search-esdb-service/proto/ml_gateway_proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCStruct struct {
	MlGatewayClient ml_gateway_proto.MlGatewayServiceClient
}

func NewMockgRPC() *GRPCStruct {
	return &GRPCStruct{
		MlGatewayClient: nil,
	}
}

func NewgRPC(cfg *config.Config) (*GRPCStruct, error) {
	mlGatewayConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.MlGateway.URL, cfg.MlGateway.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, "Error connecting to ml gateway service via gRPC "+err.Error())
	}

	mlGatewayClient := ml_gateway_proto.NewMlGatewayServiceClient(mlGatewayConn)

	return &GRPCStruct{
		MlGatewayClient: mlGatewayClient,
	}, nil

}

func (g *CommunicationImpl) Text2Vec(text string) (*ml_gateway_proto.Text2VecResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.MlGatewayClient.Text2Vec(ctx, &ml_gateway_proto.Text2VecRequest{Text: text})
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError,
			fmt.Sprintf("Error calling ml gateway service via gRPC %v", err),
		)
	}

	return result, nil
}
