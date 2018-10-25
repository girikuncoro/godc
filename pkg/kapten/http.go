package kapten

import "github.com/vmware/dispatch/pkg/http"

func httpServer(config *serverConfig) *http.Server {
	server := http.NewServer(nil)
	server.Host = config.Host
	server.Port = config.Port

	return server
}
