package gateway

import (
	"context"
	"log"
	"net/http"
	"path"
	"strings"

	pb "github.com/alkapa/quasar-fire/pkg/api/v1/quasar"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	Server interface {
		Serve() error
	}

	server struct {
		mux  *http.ServeMux
		url  string
		port string
	}

	Options struct {
		GRPCUrl         string
		RESTServicePort string
	}
)

func NewServer(opt *Options) (Server, error) {
	ctx := context.Background()
	api := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions:   protojson.MarshalOptions{},
				UnmarshalOptions: protojson.UnmarshalOptions{},
			},
		),
		runtime.WithErrorHandler(HttpError),
	)

	grpcOptions := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if err := pb.RegisterQuasarFireHandlerFromEndpoint(ctx, api, opt.GRPCUrl, grpcOptions); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)
	mux.Handle("/", api)

	return &server{
		mux:  mux,
		url:  ":" + opt.RESTServicePort,
		port: opt.RESTServicePort,
	}, nil
}

func (s *server) Serve() error {
	log.Printf("REST Listening on port: %s", s.port)
	return http.ListenAndServe(s.url, s.mux)
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("swagger-ui/", p)
	http.ServeFile(w, r, p)
}
