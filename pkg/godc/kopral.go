package main

import (
	"io"
	"net/http"

	"github.com/girikuncoro/godc/pkg/kopral"
	"github.com/girikuncoro/godc/pkg/kopral/gen/restapi"
	"github.com/girikuncoro/godc/pkg/kopral/gen/restapi/operations"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type kopralConfig struct{}

// NewCmdKopral creates a subcommand to run kopral
func NewCmdKopral(out io.Writer, config *godcConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kopral",
		Short: "Run Kopral as agent in each of GoDC node",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runKopral(config)
		},
	}
	cmd.SetOutput(out)

	return cmd
}

func runKopral(config *godcConfig) {
	kopralHandler := initKopral(config)

	server := httpServer(config)
	server.SetHandler(kopralHandler)
	defer server.Shutdown()
	if err := server.Serve(); err != nil {
		log.Error(err)
	}
}

func initKopral(config *godcConfig) http.Handler {
	swaggerSpec, err := loads.Analyzed(restapi.FlatSwaggerJSON, "2.0")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewKopralAPI(swaggerSpec)

	handlers := kopral.NewHandlers()
	handlers.ConfigureHandlers(api)

	return api.Serve(nil)
}
