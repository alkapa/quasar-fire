package server

import pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"

type (
	Satellites map[pb.SatelliteAllianceName]Satellite

	Satellite struct {
		*Point
	}

	Point struct {
		X, Y float32
	}
)
