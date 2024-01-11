package server

import (
	"auth-service/auth_proto"
	"auth-service/jwt"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	auth_proto.UnimplementedAuthServiceServer
}

func (a *GRPCServer) Authorization(ctx context.Context, req *auth_proto.AuthorizationRequest) (*auth_proto.AuthorizationResponse, error) {
	// Extract the token and requiredRole from the request
	token := req.GetToken()
	requiredRole := req.GetRequiredRole()

	tokenClaim,err := jwt.ValidateAndExtractToken(token)
	if err != nil {
		if err.StatusCode == 401 {
			return &auth_proto.AuthorizationResponse{IsAuthorized: false}, nil
		}
		return nil, err
	}

	role := tokenClaim.Role

	result := jwt.HasAuthorizeRole(role, requiredRole, true)
	
	// Return the response
	return &auth_proto.AuthorizationResponse{IsAuthorized: result}, nil
}

func GRPCListen() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()

	auth_proto.RegisterAuthServiceServer(grpcServer, &GRPCServer{})


	log.Println("gRPC server listening on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}

}
