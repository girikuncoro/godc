package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const DHCPExposerPort = "8000"

type DHCPExposer struct {
	router *httprouter.Router
}

func (d *DHCPExposer) Run() {
	d.router.GET("/health", d.handleHealth)
	d.router.GET("/ips", d.handleListIP)

	log.Fatal(http.ListenAndServe(":"+DHCPExposerPort, d.router))
}

func (d *DHCPExposer) handleHealth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "foo")
}

func (d *DHCPExposer) handleListIP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "bar")
}

func main() {
	de := &DHCPExposer{
		router: httprouter.New(),
	}

	log.Printf("Serving exposer on port " + DHCPExposerPort + " ...")
	de.Run()
}
