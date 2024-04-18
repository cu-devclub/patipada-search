package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"ml-gateway-service/config"
	"ml-gateway-service/proto"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	proto.UnimplementedSearchServiceServer
	server Server
}

func GRPCListen(server *Server, cfg *config.Config) {
	slog.Info("Starting gRPC server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterSearchServiceServer(grpcServer, &GRPCServer{server: *server})

	slog.Info("gRPC server listening on port:", slog.Int("Port", cfg.App.GRPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func (g *GRPCServer) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	slog.Info("Received search request", slog.String("Text", req.Text))
	gatewayArch := g.server.GetGatewayArch()

	response, err := gatewayArch.Usecase.Text2Vec(req.Text)
	if err != nil {
		slog.Error("Failed to search text", slog.String("Text", req.Text), slog.String("err", err.Error()))
		return nil, err
	}

	if response == nil {
		slog.Info("No result found")
		return nil,nil
	}

	protoRes := []*proto.Result{}
	for _, res := range response {
		protoRes = append(protoRes, &proto.Result{
			Name: 	 res.Name,
			Value:  res.Embedding,
			Score:  res.ScoreWeight,
		})
	}

	slog.Info("Search result", slog.String("Text", req.Text), slog.Int("Results", len(protoRes)))
	return &proto.SearchResponse{
		Results: protoRes,
	}, nil
}
