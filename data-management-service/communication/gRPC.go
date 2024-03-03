package communication

import (
	"context"
	"data-management/config"
	"data-management/proto/auth_proto"
	"data-management/proto/search_proto"
	"fmt"
	"log"
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

func NewgRPC(cfg *config.Config) *GRPCStruct {
	log.Println("Connecting to auth service via gRPC...")
	authConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.App.AuthService, cfg.App.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("Error connecting to auth service via gRPC %v", err)
		return nil
	}
	log.Println("Connected to auth service via gRPC!")

	log.Println("Connecting to search service via gRPC...")
	searchConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.App.SearchService, cfg.App.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("Error connecting to search service via gRPC %v", err)
		return nil
	}
	log.Println("Connected to search service via gRPC!")

	authClient := auth_proto.NewAuthServiceClient(authConn)
	searchClient := search_proto.NewSearchServiceClient(searchConn)

	return &GRPCStruct{
		AuthClient:   authClient,
		SearchClient: searchClient,
	}
}

func (g *CommunicationImpl) Authorization(token string, requiredRole string) (bool, error) {
	log.Println("Authorization ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.AuthClient.Authorization(ctx, &auth_proto.AuthorizationRequest{Token: token, RequiredRole: requiredRole})
	if err != nil {
		log.Println("Error calling auth service via gRPC", err)
		return false, err
	}

	if !result.IsAuthorized {
		log.Println("User is not authorized")
		return false, err
	}

	log.Println("User is authorized")
	return true, nil
}

func (g *CommunicationImpl) VerifyUsername(username string) (bool, error) {
	log.Println("Verify Username : ", username, " ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.AuthClient.VerifyUsername(ctx, &auth_proto.VerifyUsernameRequest{Username: username})
	if err != nil {
		log.Println("Error calling auth service via gRPC", err)
		return false, err
	}

	if !result.IsVerified {
		log.Println("Username is not valid")
		return false, err
	}

	log.Println("Username is valid")
	return true, nil
}

func (g *CommunicationImpl) SearchRecord(recordID string) (bool, error) {
	log.Println("Search Record : ", recordID, " ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.GRPC.SearchClient.SearchRecord(ctx, &search_proto.SearchRequest{Query: recordID})
	log.Println("SearchRecord Result", result, "ERror", err)
	if err != nil {
		if err.Error() != "rpc error: code = Unknown desc = Elasticsearch error: 405 Method Not Allowed" {
			// grpc error with elastic which can be ignored
			log.Println("Error calling search service via gRPC", err)
			return false, err
		}
	}

	if !result.IsFounded {
		log.Println("Record is not found")
		return false, err
	}

	log.Println("Record is found")
	return true, nil
}
