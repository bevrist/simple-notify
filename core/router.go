// router sets up the application router
package main

import (
	"fmt"
	"net/http"

	"github.com/bevrist/simple-notify/core/internal/receiver"

	"github.com/gorilla/mux"
)

// setupRouter creates router populated with all application endpoints
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	receiver.AddHandlers(r)

	//healthcheck
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") }).Methods(http.MethodGet, http.MethodHead)
	//default wrong method handler
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "405 wrong method.", http.StatusMethodNotAllowed)
	})
	//default not found handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "404 not found.", http.StatusNotFound)
	})
	return r
}
