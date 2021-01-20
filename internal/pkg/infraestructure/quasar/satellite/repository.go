package satellite

import (
	"sync"

	repo "github.com/alkapa/quasar-fire/internal/pkg/infraestructure/quasar/server/repository"
	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
)

var mux = sync.Mutex{}

type repository struct {
	data map[pb.SatelliteAllianceName]*pb.SatelliteSecretMessage
}

func NewRepository() repo.Repository {
	m := make(map[pb.SatelliteAllianceName]*pb.SatelliteSecretMessage)
	return &repository{data: m}
}

func (r *repository) Get(in pb.SatelliteAllianceName) (*pb.SatelliteSecretMessage, error) {
	mux.Lock()
	defer mux.Unlock()

	v, ok := r.data[in]
	if !ok {
		return nil, ErrNotFound
	}

	return v, nil
}

func (r *repository) Set(in *pb.SatelliteSecretMessage) error {
	mux.Lock()
	defer mux.Unlock()

	if _, ok := r.data[in.Name]; !ok {
		r.data[in.Name] = in
		return nil
	}

	r.data[in.Name] = in
	return nil
}

func (r *repository) GetAll() ([]*pb.SatelliteSecretMessage, error) {
	mux.Lock()
	defer mux.Unlock()

	res := make([]*pb.SatelliteSecretMessage, 0)
	for i := 1; i <= len(r.data); i++ {
		if v, ok := r.data[pb.SatelliteAllianceName(i)]; ok {
			res = append(res, v)
		}
	}

	return res, nil
}

func (r *repository) Count() int {
	return len(r.data)
}

func (r *repository) Type() string {
	return "data store in memory"
}
