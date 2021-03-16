package main

import (
	"net/http"
	"os"
	"testing"
)

// This file setup_test.go will run before the tests run

func TestMain(m *testing.M) {
	// Before running the tests do... in this function
	// then run the tests
	// then exit

	// Stop, but run m.Run() before stopping
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
