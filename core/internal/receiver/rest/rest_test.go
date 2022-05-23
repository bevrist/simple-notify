package rest

import (
	"net/http"
	"os"
	"testing"
)

// TODO: add black box style test to test rest api endpoints

// start application for black box api test
func TestMain(m *testing.M) {
	// start application for black box api tests
	go main()
	for {
		resp, _ := http.Get(pfx + "/healthz")
		if resp.StatusCode == http.StatusOK {
			break
		}
	}
	os.Exit(m.Run())
}

//TODO: run server on test
//TODO: implement black box api test
//TODO: validate entry makes it into database
