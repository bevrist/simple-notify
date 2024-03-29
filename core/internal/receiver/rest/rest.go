// rest is a http REST based message receiver
package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bevrist/simple-notify/core/internal/database"
	"github.com/bevrist/simple-notify/core/pkg/common"
)

const host string = "rest.simple-notify.com" //hostname for REST receiver endpoint

// AddHandlers returns a mux.Router with all the endpoints for the REST receiver
func AddHandlers(r *mux.Router) {
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, host+" - ok") }).Methods(http.MethodGet, http.MethodHead).Host(host)
	r.HandleFunc("/new", newMessageHandler).Methods(http.MethodPost).Host(host)
}

// newMessageHandler (/new)[POST] endpoint that accepts a message and stores it in database
func newMessageHandler(w http.ResponseWriter, r *http.Request) {
	rBody, _ := ioutil.ReadAll(r.Body)
	var message common.Message
	err := json.Unmarshal(rBody, &message)
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError) //TODO update errors to be more descriptive
		log.Println("ERROR: UpdateUserInfoHandler: " + err.Error())
		return
	}
	err = message.Validate()
	if err != nil {
		http.Error(w, "500 Bad Message.", http.StatusInternalServerError) //TODO update errors to be more descriptive
		log.Println("ERROR: UpdateUserInfoHandler: " + err.Error())
		return
	}
	err = database.NewMessage(message, "rest")
	if err != nil {
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError) //TODO update errors to be more descriptive
		log.Println("ERROR: UpdateUserInfoHandler: " + err.Error())
		return
	}
	fmt.Fprint(w, "ok")
}

// TODO add http form based POST endpoint
