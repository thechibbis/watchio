package grpc

import (
	"context"

	"github.com/thechibbis/watchio/services/indexer/internal/domain"
	"github.com/thechibbis/watchio/services/indexer/internal/usecase"
	pb "github.com/thechibbis/watchio/shared/proto/indexer"
)

type IndexerHandler struct {
	pb.UnimplementedIndexerServiceServer

	indexerUseCase *usecase.IndexerUseCase
}

func NewGRPCHandler(torrentUC *usecase.IndexerUseCase) *IndexerHandler {
	return &IndexerHandler{
		indexerUseCase: torrentUC,
	}
}

func (h *IndexerHandler) GetMovieMagnetic(ctx context.Context, req *pb.GetMovieMagneticRequest) (*pb.GetMovieMagneticResponse, error) {
	searchQuery := domain.SearchQuery{}
	searchQuery.FromGetMovieMagneticProto(req)

	torrents, err := h.indexerUseCase.GetTorrentInfo(searchQuery)
	if err != nil {
		return nil, err
	}

	torrentsPb := make([]*pb.Torrents, 0, len(torrents))

	for _, e := range torrents {
		torrent := e.ToProto()
		torrentsPb = append(torrentsPb, &torrent)
	}

	return &pb.GetMovieMagneticResponse{
		Torrents: torrentsPb,
	}, nil
}

func (h *IndexerHandler) GetTVShowMagnetic(ctx context.Context, req *pb.GetTVShowMagneticRequest) (*pb.GetTVShowMagneticResponse, error) {
	searchQuery := domain.SearchQuery{}
	searchQuery.FromGetTVShowsMagneticProto(req)

	torrents, err := h.indexerUseCase.GetTorrentInfo(searchQuery)
	if err != nil {
		return nil, err
	}

	torrentsPb := make([]*pb.Torrents, 0, len(torrents))

	for _, e := range torrents {
		torrent := e.ToProto()
		torrentsPb = append(torrentsPb, &torrent)
	}

	return &pb.GetTVShowMagneticResponse{
		Torrents: torrentsPb,
	}, nil
}
