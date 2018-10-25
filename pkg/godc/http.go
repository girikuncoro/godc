package main

import "github.com/girikuncoro/godc/pkg/http"

func httpServer(config *godcConfig) *http.Server {
	server := http.NewServer(nil)
	server.Host = config.Host
	server.Port = config.Port

	return server
}
