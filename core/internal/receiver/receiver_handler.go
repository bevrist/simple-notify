// receiver imports all receiver packages
package receiver

import (
	"github.com/gorilla/mux"

	"github.com/bevrist/simple-notify/core/internal/receiver/rest"
)

// AddHandlers adds handlers for all receivers
func AddHandlers(r *mux.Router) {
	rest.AddHandlers(r)
}
