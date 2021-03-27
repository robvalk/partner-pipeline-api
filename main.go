package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/robvalk/partner-pipeline-api/handlers"
	"github.com/robvalk/partner-pipeline-api/spec"
)

func main() {
	var port = flag.Int("port", 8081, "Port for test HTTP server")
	flag.Parse()

	swagger, err := spec.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// We now register our petStore above as the handler for the interface
	r := handlers.HandlerWithOptions(handlers.NewDeveloperHandler(), handlers.ChiServerOptions{
		Middlewares: []handlers.MiddlewareFunc{
			middleware.OapiRequestValidator(swagger),
		},
	})

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

}
