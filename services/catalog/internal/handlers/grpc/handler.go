package grpc

import (
	"context"

	"github.com/thechibbis/watchio/services/catalog/internal/domain"
	"github.com/thechibbis/watchio/services/catalog/internal/usecase"
	pb "github.com/thechibbis/watchio/shared/proto/catalog"
)

type CatalogHandler struct {
	pb.UnimplementedCatalogServiceServer

	movieUseCase   *usecase.CatalogUseCase[domain.Movie]
	tvShowsUseCase *usecase.CatalogUseCase[domain.TVShows]
}

func NewGRPCHandler(movieUC *usecase.CatalogUseCase[domain.Movie], tvShowsUC *usecase.CatalogUseCase[domain.TVShows]) *CatalogHandler {
	return &CatalogHandler{
		movieUseCase:   movieUC,
		tvShowsUseCase: tvShowsUC,
	}
}

func (h *CatalogHandler) GetTrendingMovies(ctx context.Context, req *pb.GetTrendingMoviesRequest) (*pb.GetTrendingMoviesResponse, error) {
	movies, err := h.movieUseCase.GetTrending()
	if err != nil {
		return nil, err
	}

	moviesPb := make([]*pb.Movies, 0, len(movies))

	for _, e := range movies {
		toProto := e.ToProto()
		moviesPb = append(moviesPb, &toProto)
	}

	return &pb.GetTrendingMoviesResponse{
		Movies: moviesPb,
	}, nil
}

func (h *CatalogHandler) GetTrendingTVShows(ctx context.Context, req *pb.GetTrendingTVShowsRequest) (*pb.GetTrendingTVShowsResponse, error) {
	tvShows, err := h.tvShowsUseCase.GetTrending()
	if err != nil {
		return nil, err
	}

	tvShowsPb := make([]*pb.TVShows, 0, len(tvShows))

	for _, e := range tvShows {
		toProto := e.ToProto()
		tvShowsPb = append(tvShowsPb, &toProto)
	}

	return &pb.GetTrendingTVShowsResponse{
		TvShows: tvShowsPb,
	}, nil
}
