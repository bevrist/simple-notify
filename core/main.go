// Frontend-api provides a public API for clients interacting with the app
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bevrist/simple-notify/core/internal/receiver"

	"github.com/gorilla/mux"
)

// Global Variables
var apiVersion string = "1.0" //the api version this service implements

func main() {
	//specify routes and start http server
	r := mux.NewRouter()
	receiver.AddHandlers(r)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })
	var handlers http.Handler = r
	log.Println("Frontend-api listening at address 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handlers))
}
