package jackett

import (
	"context"
	"os"

	"github.com/thechibbis/watchio/services/indexer/internal/domain"
	"github.com/webtor-io/go-jackett"
)

type JackettRepository struct {
	Client *jackett.Jackett
}

func NewJackettRepository() *JackettRepository {
	j := jackett.NewJackett(&jackett.Settings{
		ApiURL: os.Getenv("JACKETT_API_URL"),
		ApiKey: os.Getenv("JACKETT_API_KEY"),
	})

	return &JackettRepository{j}
}

func (r *JackettRepository) GetTorrentInfo(searchQuery domain.SearchQuery) ([]domain.TorrentInfo, error) {
	ctx := context.Background()

	query := searchQuery.QueryString()

	resp, err := r.Client.Fetch(ctx, &jackett.FetchRequest{
		Query: query,
	})

	if err != nil {
		return nil, err
	}

	parsedResults := make([]domain.TorrentInfo, 0, len(resp.Results))

	// TODO! do smth with the release date from catalog service

	for _, e := range resp.Results {
		if e.MagnetUri == "" {
			continue
		}

		parsedResults = append(parsedResults, domain.TorrentInfo{
			MagneticURI: e.MagnetUri,
			Size:        e.Size,
			Seeders:     e.Seeders,
			Tracker:     e.Tracker,
		})
	}

	return parsedResults, nil
}
