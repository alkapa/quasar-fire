package server

import "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"

var allianceSatellites = Satellites{
	quasar.SatelliteAllianceName_kenobi: {
		Point: &Point{
			X: -500,
			Y: -200,
		},
	},
	quasar.SatelliteAllianceName_skywalker: {
		Point: &Point{
			X: 100,
			Y: -100,
		},
	},
	quasar.SatelliteAllianceName_sato: {
		Point: &Point{
			X: 500,
			Y: 100,
		},
	},
}
