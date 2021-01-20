package repository

import pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"

type Repository interface {
	Get(id pb.SatelliteAllianceName) (*pb.SatelliteSecretMessage, error)
	Set(in *pb.SatelliteSecretMessage) error
	GetAll() ([]*pb.SatelliteSecretMessage, error)
	Count() int
	Type() string
}
