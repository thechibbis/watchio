package usecase

import "github.com/thechibbis/watchio/services/indexer/internal/domain"

type IndexerRepository interface {
	GetTorrentInfo(searchQuery domain.SearchQuery) ([]domain.TorrentInfo, error)
}
