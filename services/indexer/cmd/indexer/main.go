package main

import (
	"log"
	"net"

	grpc_handler "github.com/thechibbis/watchio/services/indexer/internal/handlers/grpc"
	"github.com/thechibbis/watchio/services/indexer/internal/repository/jackett"
	"github.com/thechibbis/watchio/services/indexer/internal/usecase"
	pb "github.com/thechibbis/watchio/shared/proto/indexer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":50053"

	indexerRepo := jackett.NewJackettRepository()

	indexerUseCase := usecase.NewIndexerUseCase(indexerRepo)

	indexerHandler := grpc_handler.NewGRPCHandler(indexerUseCase)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen at: %s\n%v", addr, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterIndexerServiceServer(grpcServer, indexerHandler)
	reflection.Register(grpcServer)

	log.Printf("Indexer service listening at: %s\n", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize indexer service: %v", err)
	}
}
