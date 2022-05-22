// receiver imports all receiver packages
package receiver

import (
	"github.com/gorilla/mux"

	rest "github.com/bevrist/simple-notify/core/internal/receiver/rest"
)

// AddHandlers returns a mux.Router with all the endpoints for each receiver
func AddHandlers(r *mux.Router) {
	rest.AddHandlers(r)
}
