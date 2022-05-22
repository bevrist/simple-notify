// rest is a http REST based message receiver
package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: expose endpoint that accepts user token and message
// TODO: return error codes as necessary
// TODO: store message in database

// AddHandlers returns a mux.Router with all the endpoints for the REST receiver
func AddHandlers(r *mux.Router) {
	r.HandleFunc("/rest/test", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "test") }).Methods(http.MethodGet, http.MethodHead)
}
