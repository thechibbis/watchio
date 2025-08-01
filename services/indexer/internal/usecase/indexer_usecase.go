package usecase

import "github.com/thechibbis/watchio/services/indexer/internal/domain"

type IndexerUseCase struct {
	repo IndexerRepository
}

func NewIndexerUseCase(r IndexerRepository) *IndexerUseCase {
	return &IndexerUseCase{repo: r}
}

func (uc *IndexerUseCase) GetTorrentInfo(searchQuery domain.SearchQuery) ([]domain.TorrentInfo, error) {
	return uc.repo.GetTorrentInfo(searchQuery)
}
