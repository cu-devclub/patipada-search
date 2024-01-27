package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"search-esdb-service/config"
	search_proto "search-esdb-service/proto/search_proto"
	"search-esdb-service/record/models"
	recordRepository "search-esdb-service/record/repositories"
	recordUsecases "search-esdb-service/record/usecases"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	search_proto.UnimplementedSearchServiceServer
	server Server
}

func GRPCListen(server Server, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()

	search_proto.RegisterSearchServiceServer(grpcServer, &GRPCServer{server: server})

	log.Println("gRPC server listening on port:", cfg.App.GRPCPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func (a *GRPCServer) SearchRecord(ctx context.Context, req *search_proto.SearchRequest) (*search_proto.SearchResponse, error) {
	recordESRepository := recordRepository.NewRecordESRepository(a.server.GetDB())

	recordUsecase := recordUsecases.NewRecordUsecase(recordESRepository)

	result, err := recordUsecase.SearchByRecordIndex("record", req.Query)
	if err != nil {
		log.Println("RERERERERE", err)
		return nil, err
	}

	if result != nil {
		log.Println("search result", result.ToString())
	}

	log.Println("Reteruwfuisduifgjsdiuof")
	return &search_proto.SearchResponse{
		IsFounded: result != nil,
	}, nil
}

func (a *GRPCServer) UpdateRecord(ctx context.Context, req *search_proto.UpdateRecordRequest) (*search_proto.UpdateRecordResponse, error) {
	log.Println("Receiving update record request from gRPC client...")
	recordESRepository := recordRepository.NewRecordESRepository(a.server.GetDB())

	recordUsecase := recordUsecases.NewRecordUsecase(recordESRepository)

	record := &models.UpdateRecord{
		DocumentID: req.Index,
		Question:   req.Question,
		Answer:     req.Answer,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	}

	err := recordUsecase.UpdateRecord(record)
	if err != nil {
		return &search_proto.UpdateRecordResponse{
			IsUpdated: false,
		}, err
	}

	return &search_proto.UpdateRecordResponse{
		IsUpdated: true,
	}, nil
}