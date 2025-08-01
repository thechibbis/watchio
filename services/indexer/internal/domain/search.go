package domain

import (
	"fmt"
	"strings"

	pb "github.com/thechibbis/watchio/shared/proto/indexer"
)

type SearchQuery struct {
	Title       string
	ReleaseDate string
	Episode     int32
	Season      int32
}

func (s *SearchQuery) QueryString() string {
	var builder strings.Builder

	formattedTitle := fmt.Sprintf("\"%s\"", s.Title)

	builder.WriteString(formattedTitle)

	if s.Season > 0 && s.Episode > 0 {
		formattedSeason := fmt.Sprintf("%02d", s.Season)
		formattedEpisode := fmt.Sprintf("%02d", s.Episode)

		fmt.Fprintf(&builder, " S%sE%s", formattedSeason, formattedEpisode)
	}

	return builder.String()
}

func (s *SearchQuery) FromGetMovieMagneticProto(proto *pb.GetMovieMagneticRequest) {
	s.Title = proto.Name
	s.ReleaseDate = proto.ReleaseDate
	s.Episode = 0
	s.Season = 0
}

func (s *SearchQuery) FromGetTVShowsMagneticProto(proto *pb.GetTVShowMagneticRequest) {
	s.Title = proto.Name
	s.ReleaseDate = proto.ReleaseDate
	s.Episode = proto.Episode
	s.Season = proto.Season
}
