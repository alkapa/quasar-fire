package server

import (
	"context"
	"errors"
	"math"
	"strings"

	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
)

func GetLocation(distances ...float32) (x, y float32) {
	point := getLocation(
		allianceSatellites[pb.SatelliteAllianceName_kenobi].Point,
		allianceSatellites[pb.SatelliteAllianceName_skywalker].Point,
		allianceSatellites[pb.SatelliteAllianceName_sato].Point,
		distances[pb.SatelliteAllianceName_kenobi-1],
		distances[pb.SatelliteAllianceName_skywalker-1],
		distances[pb.SatelliteAllianceName_sato-1],
	)
	return point.X, point.Y
}

func GetMessage(messages ...[]string) (msg string) {

	msgCount := len(messages)
	if !(msgCount > 0) {
		return ""
	}

	msgContentCount := len(messages[0])

	{
		msg := make([]string, 0)
		for j := 0; j < msgContentCount; j++ {
			msgMap := make(map[string]int)
			for _, msgList := range messages {
				msg := msgList[j]
				if msg != "" {
					msgMap[msg] = j
				}
			}

			if len(msgMap) > 1 {
				return ""
			}

			for v := range msgMap {
				msg = append(msg, v)
			}
		}

		if len(msg) != msgContentCount {
			return ""
		}

		return strings.Join(msg, " ")
	}
}

func getLocation(p1, p2, p3 *Point, d1, d2, d3 float32) *Point {

	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y
	x3, y3 := p3.X, p3.Y

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := (d1 * d1) - (d2 * d2) - (x1 * x1) + (x2 * x2) - (y1 * y1) + (y2 * y2)

	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := (d2 * d2) - (d3 * d3) - (x2 * x2) + (x3 * x3) - (y2 * y2) + (y3 * y3)

	X := E*A - B*D
	Y := B*D - A*E

	x := (C*E - F*B) / X
	y := (C*D - A*F) / Y

	return &Point{x, y}
}

func topSecretResolver(_ context.Context, in *pb.TopSecretRequest) (*pb.TopSecretResponse, error) {
	if !(len(in.Satellites) == 3) {
		return nil, errors.New("message quantity not valid")
	}

	satelliteMap := make(map[pb.SatelliteAllianceName]*pb.SatelliteSecretMessage)
	for _, v := range in.Satellites {
		if _, ok := satelliteMap[v.Name]; ok {
			return nil, errors.New("satellite alliance name it`s duplicate")
		}
		satelliteMap[v.Name] = v
	}

	distance := make([]float32, 0, 3)
	messages := make([][]string, 0)
	messageLen := len(satelliteMap[pb.SatelliteAllianceName(1)].Message)

	for i := 1; i <= 3; i++ {
		v := satelliteMap[pb.SatelliteAllianceName(i)]

		distance = append(distance, v.Distance)

		if !(messageLen == len(v.Message)) {
			return nil, errors.New("messages len not equal")
		}
		messages = append(messages, v.Message)
	}

	x, y := GetLocation(distance...)
	if math.IsInf(float64(x), 0) || math.IsInf(float64(y), 0) {
		return nil, errors.New("can`t get localization")
	}

	msg := GetMessage(messages...)
	if msg == "" {
		return nil, errors.New("can`t get message")
	}

	return &pb.TopSecretResponse{
		Position: &pb.Position{
			X: x,
			Y: y,
		},
		Message: msg,
	}, nil
}
