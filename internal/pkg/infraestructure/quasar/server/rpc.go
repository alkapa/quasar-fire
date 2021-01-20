package server

import (
	"context"

	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) TopSecret(ctx context.Context, in *pb.TopSecretRequest) (*pb.TopSecretResponse, error) {
	log := s.log.WithFields(
		logrus.Fields{
			"rpc": "TopSecret",
		},
	)

	res, err := topSecretResolver(ctx, in)
	if err != nil {
		log.WithFields(
			logrus.Fields{
				"payload": in,
				"from":    "request",
			},
		).WithError(err).Error("top secret resolver error")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	for _, v := range in.Satellites {
		if err := s.repository.Set(v); err != nil {
			msg := "can`t store data"
			log.WithFields(
				logrus.Fields{
					"repository": s.repository.Type(),
					"action":     "Set",
					"input":      v,
				},
			).WithError(err).Error(msg)
			return nil, status.Error(codes.DataLoss, msg)
		}
	}

	log.Info("successful")

	return res, nil
}

func (s *server) TopSecretSplitGet(ctx context.Context, _ *any.Any) (*pb.TopSecretResponse, error) {
	log := s.log.WithFields(
		logrus.Fields{
			"rpc": "TopSecretSplitGet",
		},
	)

	if !(s.repository.Count() > 0) {
		msg := "not sufficient data"
		log.WithFields(
			logrus.Fields{
				"repository": s.repository.Type(),
				"dataCount":  s.repository.Count(),
			},
		).Error(msg)
		return nil, status.Error(codes.NotFound, msg)
	}

	list, err := s.repository.GetAll()
	if err != nil {
		msg := "can`t get stored data"
		log.WithFields(
			logrus.Fields{
				"repository": s.repository.Type(),
				"action":     "GetAll",
				"dataCount":  s.repository.Count(),
			},
		).WithError(err).Error(msg)
		return nil, status.Error(codes.DataLoss, msg)
	}

	res, err := topSecretResolver(ctx,
		&pb.TopSecretRequest{
			Satellites: list,
		},
	)
	if err != nil {
		log.WithFields(
			logrus.Fields{
				"action": "topSecretResolver",
				"from":   "repository",
				"input":  list,
			},
		).WithError(err).Error("top secret resolver error")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	log.Info("successful")

	return res, nil
}

func (s *server) TopSecretSplitSet(ctx context.Context, in *pb.SatelliteSecretMessage) (*pb.TopSecretResponse, error) {
	log := s.log.WithFields(
		logrus.Fields{
			"rpc": "TopSecretSplitSet",
		},
	)

	if err := s.repository.Set(in); err != nil {
		msg := "can`t set data"
		log.WithFields(
			logrus.Fields{
				"repository": s.repository.Type(),
				"action":     "Set",
				"input":      in,
			},
		).WithError(err).Error(msg)
		return nil, status.Error(codes.DataLoss, msg)
	}

	res, err := s.TopSecretSplitGet(ctx, &any.Any{})
	if err != nil {
		return nil, err
	}

	log.Info("successful")

	return res, nil
}
