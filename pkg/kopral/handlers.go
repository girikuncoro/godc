package kopral

import (
	"log"

	"github.com/girikuncoro/godc/pkg/kopral/gen/models"
	"github.com/girikuncoro/godc/pkg/kopral/gen/restapi/operations"
	"github.com/girikuncoro/godc/pkg/kopral/gen/restapi/operations/kopral"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// KopralFlags are configuration flags for Kopral
var KopralFlags = struct {
	Config string `long:"config" description:"Path to Config file" default:"./config.dev.json"`
}{}

// Handlers define a set of handlers for Kopral
type Handlers struct{}

// NewHandlers create a new Kopral Handler
func NewHandlers() *Handlers {
	return &Handlers{}
}

// ConfigureHandlers configure handlers for Kopral
func (h *Handlers) ConfigureHandlers(routableAPI middleware.RoutableAPI) {
	k, ok := routableAPI.(*operations.KopralAPI)
	if !ok {
		panic("Cannot configure Kopral API")
	}

	k.Logger = log.Printf
	k.KopralGetNodeStatsHandler = kopral.GetNodeStatsHandlerFunc(h.getNodeStats)
}

func (h *Handlers) getNodeStats(params kopral.GetNodeStatsParams) middleware.Responder {
	// TODO(giri): collect proper node stats
	nodeStats := &models.NodeStats{
		CPU:    swag.Int64(20),
		Memory: swag.Int64(10),
	}
	return kopral.NewGetNodeStatsOK().WithPayload(nodeStats)
}
