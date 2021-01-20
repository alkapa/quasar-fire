package quasar

import (
	"log"
	"net"
	"os"

	"github.com/alkapa/quasar-fire/internal/pkg/infraestructure/quasar/gateway"
	repository "github.com/alkapa/quasar-fire/internal/pkg/infraestructure/quasar/satellite"
	"github.com/alkapa/quasar-fire/internal/pkg/infraestructure/quasar/server"
	"github.com/spf13/cobra"
)

var Serve = &cobra.Command{
	Use:   "quasar",
	Short: "serve quasar",
	Long:  "Servicio de rastreo y comunicacion de la alianza",
	RunE:  serve,
}

func serve(_ *cobra.Command, _ []string) error {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}

	grpcSrv, err := server.NewServer(
		&server.Options{
			Repository: repository.NewRepository(),
		},
	)
	if err != nil {
		return err
	}

	go func() {
		log.Fatal(grpcSrv.Serve(lis))
	}()

	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		restPort = "8080"
	}

	restSrv, err := gateway.NewServer(
		&gateway.Options{
			GRPCUrl:         ":" + grpcPort,
			RESTServicePort: restPort,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("listening on grpc port: %s", grpcPort)
	return restSrv.Serve()
}
