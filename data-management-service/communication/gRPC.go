package communication

import (
	"context"
	"data-management/config"
	"data-management/errors"
	"data-management/messages"
	"data-management/proto/auth_proto"
	"data-management/proto/search_proto"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCStruct struct {
	AuthClient   auth_proto.AuthServiceClient
	SearchClient search_proto.SearchServiceClient
}

func NewMockgRPC() *GRPCStruct {
	return &GRPCStruct{
		AuthClient:   nil,
		SearchClient: nil,
	}
}

func NewgRPC(cfg *config.Config) (*GRPCStruct, error) {
	authConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.App.AuthService, cfg.App.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, errors.CreateError(500, "Error connecting to auth service via gRPC "+err.Error())
	}

	searchConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.App.SearchService, cfg.App.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, errors.CreateError(500, "Error connecting to search service via gRPC "+err.Error())
	}

	authClient := auth_proto.NewAuthServiceClient(authConn)
	searchClient := search_proto.NewSearchServiceClient(searchConn)

	return &GRPCStruct{
		AuthClient:   authClient,
		SearchClient: searchClient,
	}, nil
}

func (g *CommunicationImpl) Authorization(token string, requiredRole string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.AuthClient.Authorization(ctx, &auth_proto.AuthorizationRequest{Token: token, RequiredRole: requiredRole})
	if err != nil {
		return false, errors.CreateError(http.StatusInternalServerError,
			fmt.Sprintf("Error calling auth service via gRPC %v", err),
		)
	}

	if !result.IsAuthorized {
		return false, errors.CreateError(http.StatusForbidden, "User is not authorized")
	}

	return true, nil
}

func (g *CommunicationImpl) VerifyUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.AuthClient.VerifyUsername(ctx, &auth_proto.VerifyUsernameRequest{Username: username})
	if err != nil {
		return false, errors.CreateError(http.StatusInternalServerError,
			fmt.Sprintf("Error calling auth service via gRPC %v", err),
		)
	}

	if !result.IsVerified {
		return false, errors.CreateError(http.StatusForbidden, "Username is invalid")
	}

	return true, nil
}

func (g *CommunicationImpl) SearchRecord(recordID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.SearchClient.SearchRecord(ctx, &search_proto.SearchRequest{Query: recordID})
	if err != nil && err.Error() != messages.ELASTIC_METHOD_NOT_ALLOW {
		return false, errors.CreateError(http.StatusInternalServerError,
			fmt.Sprintf("Error calling search service via gRPC %v", err),
		)
	}

	if !result.IsFounded {
		return false, errors.CreateError(http.StatusNotFound, messages.RECORD_NOT_FOUND)
	}

	return true, nil
}
