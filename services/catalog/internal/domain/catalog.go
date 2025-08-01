package domain

import pb "github.com/thechibbis/watchio/shared/proto/catalog"

type Movie struct {
	ID          int64
	Title       string
	GenreIds    []int64
	ReleaseDate string
}

func (m *Movie) ToProto() pb.Movies {
	return pb.Movies{
		Id:          m.ID,
		Title:       m.Title,
		GenreIds:    m.GenreIds,
		ReleaseDate: m.ReleaseDate,
	}
}

type TVShows struct {
	ID          int64
	Title       string
	GenreIds    []int64
	ReleaseDate string
}

func (m *TVShows) ToProto() pb.TVShows {
	return pb.TVShows{
		Id:          m.ID,
		Title:       m.Title,
		GenreIds:    m.GenreIds,
		ReleaseDate: m.ReleaseDate,
	}
}
