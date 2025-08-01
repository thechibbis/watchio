package domain

import pb "github.com/thechibbis/watchio/shared/proto/indexer"

type TorrentInfo struct {
	MagneticURI string
	Size        uint
	Seeders     uint
	Tracker     string
}

func (t *TorrentInfo) ToProto() pb.Torrents {
	return pb.Torrents{
		MagneticUri: t.MagneticURI,
		Size:        uint32(t.Size),
		Seeders:     uint32(t.Seeders),
		Tracker:     t.Tracker,
	}
}
