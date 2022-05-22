// Frontend-api provides a public API for clients interacting with the app
package main

import (
	"log"
	"net/http"
)

func main() {
	var handlers http.Handler = setupRouter()
	log.Println("listening at address 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handlers))
}
