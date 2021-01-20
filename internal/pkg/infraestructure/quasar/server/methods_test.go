package server

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
)

func TestGetMessage(t *testing.T) {
	type args struct {
		messages [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		{
			name: "",
			args: args{
				[][]string{
					{"este", "", "", "mensaje", ""},
					{"", "es", "", "", "secreto"},
					{"este", "", "un", "", ""},
				},
			},
			wantMsg: "este es un mensaje secreto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := GetMessage(tt.args.messages...); gotMsg != tt.wantMsg {
				t.Errorf("GetMessage() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	type args struct {
		distances []float32
	}
	tests := []struct {
		name  string
		args  args
		wantX float32
		wantY float32
	}{
		{
			name: "get position pdf example",
			args: args{
				distances: []float32{
					100,   //kenobi
					115.5, //skywalker
					142.7, //sato
				},
			},
			wantX: -487.2859,
			wantY: 1557.0142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := GetLocation(tt.args.distances...)
			if gotX != tt.wantX {
				t.Errorf("GetLocation() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetLocation() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func Test_topSecretResolver(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.TopSecretRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.TopSecretResponse
		wantErr bool
	}{
		{
			name: "test request pdf example",
			args: args{
				ctx: context.Background(),
				in: &pb.TopSecretRequest{
					Satellites: []*pb.SatelliteSecretMessage{
						{
							Name:     pb.SatelliteAllianceName_kenobi,
							Distance: 100,
							Message:  []string{"este", "", "", "mensaje", ""},
						},
						{
							Name:     pb.SatelliteAllianceName_skywalker,
							Distance: 115.5,
							Message:  []string{"", "es", "", "", "secreto"},
						},
						{
							Name:     pb.SatelliteAllianceName_sato,
							Distance: 142.7,
							Message:  []string{"este", "", "un", "", ""},
						},
					},
				},
			},
			want: &pb.TopSecretResponse{
				Position: &pb.Position{
					X: -487.2859,
					Y: 1557.0142,
				},
				Message: "este es un mensaje secreto",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := topSecretResolver(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("topSecretResolver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topSecretResolver() got = %v, want %v", got, tt.want)
			}
		})
	}
}
