package http

import (
	"net/http"
	"strings"
)

// AllInOneRouter implements a simple HTTP handler that routes requests to proper sub-service handlers
// When executing GoDC in a single binary mode.
type AllInOneRouter struct {
	KopralHandler http.Handler
}

// ServeHTTP implements the http.Handler interface
func (d *AllInOneRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	components := strings.SplitN(path[1:], "/", 3)
	if len(components) < 2 {
		rw.Header().Add("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(nil)
		return
	}
	// version
	if components[0] != "v1" {
		rw.Header().Add("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(nil)
		return
	}
	switch components[1] {
	case "kopral":
		d.KopralHandler.ServeHTTP(rw, r)
	default:
		rw.Header().Add("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(nil)
	}
}
