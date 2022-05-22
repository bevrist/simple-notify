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
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") }).Methods(http.MethodGet, http.MethodHead)
	return r
}
