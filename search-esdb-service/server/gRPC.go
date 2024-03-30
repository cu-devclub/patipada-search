package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"search-esdb-service/config"
	search_proto "search-esdb-service/proto/search_proto"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	search_proto.UnimplementedSearchServiceServer
	server Server
}

func GRPCListen(server Server, cfg *config.Config) {
	slog.Info("Starting gRPC server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()

	search_proto.RegisterSearchServiceServer(grpcServer, &GRPCServer{server: server})

	slog.Info("gRPC server listening on port:", slog.Int("Port", cfg.App.GRPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func (a *GRPCServer) SearchRecord(ctx context.Context, req *search_proto.SearchRequest) (*search_proto.SearchResponse, error) {
	slog.Info("SearchRecord called with:", slog.String("Query", req.Query))
	recordArch := a.server.GetRecordArch()
	result, err := recordArch.Usecase.SearchByRecordIndex("record", req.Query)

	if err != nil {
		slog.Error("Failed to search record",
			slog.String("Query", req.Query),
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	if result != nil {
		slog.Info("Record found", slog.String("Record", result.Index))
	}

	return &search_proto.SearchResponse{
		IsFounded: result != nil,
	}, nil
}
