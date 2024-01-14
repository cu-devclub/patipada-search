package communication

import (
	"context"
	"data-management/config"
	"data-management/proto/auth_proto"
	"data-management/proto/search_proto"
	"data-management/request/entities"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPC struct {
	authClient   auth_proto.AuthServiceClient
	searchClient search_proto.SearchServiceClient
}

// As a test for other feature that doesn't require gRPC,
// we can create a temporary gRPC client that does nothing.
func NewTempgRPC() Communication {
	return &gRPC{}
}

func NewgRPC(cfg *config.Config) Communication {
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

	return &gRPC{
		authClient:   authClient,
		searchClient: searchClient,
	}
}

func (g *gRPC) Authorization(token string, requiredRole string) (bool, error) {
	log.Println("Authorization ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.authClient.Authorization(ctx, &auth_proto.AuthorizationRequest{Token: token, RequiredRole: requiredRole})
	if err != nil {
		log.Println("Error calling auth service via gRPC", err)
		return false, err
	}

	if result.IsAuthorized != true {
		log.Println("User is not authorized")
		return false, err
	}

	log.Println("User is authorized")
	return true, nil
}

func (g *gRPC) VerifyUsername(username string) (bool, error) {
	log.Println("Verify Username : ", username, " ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.authClient.VerifyUsername(ctx, &auth_proto.VerifyUsernameRequest{Username: username})
	if err != nil {
		log.Println("Error calling auth service via gRPC", err)
		return false, err
	}

	if result.IsVerified != true {
		log.Println("Username is not valid")
		return false, err
	}

	log.Println("Username is valid")
	return true, nil
}

func (g *gRPC) SearchRecord(recordID string) (bool, error) {
	log.Println("Search Record : ", recordID, " ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.searchClient.SearchRecord(ctx, &search_proto.SearchRequest{Query: recordID})
	log.Println("SearchRecord Result", result,"ERror",err)
	if err != nil  {
		if err.Error() != "rpc error: code = Unknown desc = Elasticsearch error: 405 Method Not Allowed" {
			// grpc error with elastic which can be ignored
			log.Println("Error calling search service via gRPC", err)
			return false, err
		}
	}

	if result.IsFounded != true {
		log.Println("Record is not found")
		return false, err
	}

	log.Println("Record is found")
	return true, nil
}

func (g *gRPC) UpdateRecord(record *entities.Record) (bool, error) {
	log.Println("Update Record : ", record, " ....")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := g.searchClient.UpdateRecord(ctx, &search_proto.UpdateRecordRequest{
		Index:     record.Index,
		Question:  record.Question,
		Answer:    record.Answer,
		StartTime: record.StartTime,
		EndTime:   record.EndTime,
	})
	if err != nil {
		log.Println("Error calling search service via gRPC", err)
		return false, err
	}

	if result.IsUpdated != true {
		log.Println("Record is not updated")
		return false, nil
	}

	log.Println("Record is updated")
	return true, nil
}
