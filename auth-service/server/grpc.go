package server

import (
	"auth-service/auth_proto"
	"auth-service/config"
	"auth-service/jwt"
	usersRepositories "auth-service/users/repositories"
	usersUsecases "auth-service/users/usecases"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	auth_proto.UnimplementedAuthServiceServer
	server Server
}

func GRPCListen(server Server, cfg *config.Config) {
	log.Println("Starting gRPC server....")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()

	auth_proto.RegisterAuthServiceServer(grpcServer, &GRPCServer{server: server})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}

	log.Println("gRPC server listening on port:", cfg.App.GRPCPort)

}

func (a *GRPCServer) Authorization(ctx context.Context, req *auth_proto.AuthorizationRequest) (*auth_proto.AuthorizationResponse, error) {
	slog.Info("Recieving gRPC connection for authorization....")
	// Extract the token and requiredRole from the request
	token := req.GetToken()
	requiredRole := req.GetRequiredRole()

	tokenClaim, err := jwt.ValidateAndExtractToken(token)
	if err != nil {
		slog.Error("Error while validating and extracting token: ",
			slog.Any("error", err),
		)
		if err.StatusCode == 401 {
			return &auth_proto.AuthorizationResponse{IsAuthorized: false}, nil
		}
		return nil, err
	}

	role := tokenClaim.Role
	result := jwt.HasAuthorizeRole(role, requiredRole, true)

	return &auth_proto.AuthorizationResponse{IsAuthorized: result}, nil
}

func (a *GRPCServer) VerifyUsername(ctx context.Context, req *auth_proto.VerifyUsernameRequest) (*auth_proto.VerifyUsernameResponse, error) {
	slog.Info("Recieving gRPC connection for verifying username....")
	username := req.GetUsername()
	usersPostgresRepository := usersRepositories.NewUsersPostgresRepository(a.server.GetDB())
	usersUsecase := usersUsecases.NewUsersUsecaseImpl(
		usersPostgresRepository,
		nil,
	)

	result, err := usersUsecase.VerifyUsername(username)
	if err != nil {
		slog.Error("Error while verifying username: ",
			slog.Any("error", err),
		)
		return nil, err
	}

	return &auth_proto.VerifyUsernameResponse{IsVerified: result}, nil
}
