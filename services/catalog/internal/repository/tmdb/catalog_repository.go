package tmdb

import (
	"fmt"
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/thechibbis/watchio/services/catalog/internal/domain"
)

type TMDBRepository[T domain.Movie | domain.TVShows] struct {
	Client *tmdb.Client
}

func NewTMBDRepository[T domain.Movie | domain.TVShows]() *TMDBRepository[T] {
	client, err := tmdb.Init(os.Getenv("TMDB_API_KEY"))
	if err != nil {
		log.Fatalf("Error creating a tmdb client: %s", err.Error())
	}

	client.SetClientAutoRetry()

	return &TMDBRepository[T]{Client: client}
}

func (r *TMDBRepository[T]) GetTrending() ([]T, error) {
	var mediaType string
	var zero T

	switch any(zero).(type) {
	case domain.Movie:
		mediaType = "movie"
	case domain.TVShows:
		mediaType = "tv"
	default:
		return nil, fmt.Errorf("Media type not supported: %T", zero)
	}

	trending, err := r.Client.GetTrending(mediaType, "week", map[string]string{})
	if err != nil {
		return nil, err
	}

	parsedTrending := make([]T, 0, len(trending.Results))

	for _, e := range trending.Results {
		title := e.Title
		releaseDate := e.ReleaseDate

		if mediaType == "tv" {
			title = e.Name
			releaseDate = e.FirstAirDate
		}

		parsedTrending = append(parsedTrending, T{
			ID:          e.ID,
			Title:       title,
			GenreIds:    e.GenreIDs,
			ReleaseDate: releaseDate,
		})
	}

	return parsedTrending, nil
}
