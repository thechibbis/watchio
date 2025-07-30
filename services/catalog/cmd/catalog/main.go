package main

import (
	"log"
	"net"

	"github.com/thechibbis/watchio/services/catalog/internal/domain"
	grpc_handler "github.com/thechibbis/watchio/services/catalog/internal/handlers/grpc"
	"github.com/thechibbis/watchio/services/catalog/internal/repository/tmdb"
	"github.com/thechibbis/watchio/services/catalog/internal/usecase"
	pb "github.com/thechibbis/watchio/shared/proto/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":50052"

	movieRepo := tmdb.NewTMBDRepository[domain.Movie]()
	tvShowsRepo := tmdb.NewTMBDRepository[domain.TVShows]()

	movieUseCase := usecase.NewCatalogUseCase(movieRepo)
	tvShowsUseCase := usecase.NewCatalogUseCase(tvShowsRepo)

	catalogHandler := grpc_handler.NewGRPCHandler(movieUseCase, tvShowsUseCase)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen at: %s\n%v", addr, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCatalogServiceServer(grpcServer, catalogHandler)
	reflection.Register(grpcServer)

	log.Printf("Catalog Service running at: %s\n", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot initialize catalog service: %v\n", err)
	}
}
