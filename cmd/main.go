package main

import (
	"github.com/alkapa/quasar-fire/cmd/quasar"
	"github.com/spf13/cobra"
)

var (
	cmd   = &cobra.Command{}
	serve = &cobra.Command{
		Use: "serve",
	}
)

func init() {
	cmd.AddCommand(serve)
	{
		serve.AddCommand(quasar.Serve)
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
