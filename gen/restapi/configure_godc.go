// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/girikuncoro/godc/gen/restapi/operations"
	"github.com/girikuncoro/godc/gen/restapi/operations/cluster"
	"github.com/girikuncoro/godc/gen/restapi/operations/vm"
)

//go:generate swagger generate server --target .. --name godc --spec ../../swagger/swagger.yaml --client-package godc

func configureFlags(api *operations.GodcAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GodcAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ClusterGetClusterHandler = cluster.GetClusterHandlerFunc(func(params cluster.GetClusterParams) middleware.Responder {
		return middleware.NotImplemented("operation cluster.GetCluster has not yet been implemented")
	})
	api.VMGetVMHandler = vm.GetVMHandlerFunc(func(params vm.GetVMParams) middleware.Responder {
		return middleware.NotImplemented("operation vm.GetVM has not yet been implemented")
	})
	api.ClusterListClustersHandler = cluster.ListClustersHandlerFunc(func(params cluster.ListClustersParams) middleware.Responder {
		return middleware.NotImplemented("operation cluster.ListClusters has not yet been implemented")
	})
	api.VMListVmsHandler = vm.ListVmsHandlerFunc(func(params vm.ListVmsParams) middleware.Responder {
		return middleware.NotImplemented("operation vm.ListVms has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
