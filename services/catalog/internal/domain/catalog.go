package domain

type Movie struct {
	ID          int64
	Title       string
	GenreIds    []int64
	ReleaseDate string
}

type TVShows struct {
	ID          int64
	Title       string
	GenreIds    []int64
	ReleaseDate string
}
