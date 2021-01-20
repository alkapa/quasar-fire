package server

import (
	repo "github.com/alkapa/quasar-fire/internal/pkg/infraestructure/quasar/server/repository"
	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
	"github.com/alkapa/quasar-fire/utils/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type (
	Options struct {
		Repository repo.Repository
	}

	server struct {
		repository repo.Repository
		log        *logrus.Entry
	}
)

func NewServer(opt *Options) (*grpc.Server, error) {
	srv, err := newServer(opt)
	if err != nil {
		return nil, err
	}

	srvGrpc := grpc.NewServer()

	pb.RegisterQuasarFireServer(srvGrpc, srv)
	return srvGrpc, nil
}

func newServer(opt *Options) (pb.QuasarFireServer, error) {
	return &server{
		repository: opt.Repository,
		log: logger.New().WithFields(
			logrus.Fields{
				"service": "quasar-fire",
				"version": "1",
			},
		),
	}, nil
}
