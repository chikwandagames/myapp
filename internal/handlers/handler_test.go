package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Instead of writing single test, we can write table tests

// postData is For form data, eg posting to the make-reservation page
type postData struct {
	key   string
	value string
}

// For running multiple tests
var theTests = []struct {
	// What we want to call an individual test
	name string
	// Path matche by the routes
	url string
	// Post or Get
	method string
	// Things being posted
	params []postData
	// Test to see what response code we are getting, eg 200
	expectedStatusCode int
}{
	// Define and populate with data all at once
	// Because we have a slice of structs
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"general", "/generals", "GET", []postData{}, http.StatusOK},
	{"major", "/majors", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	// getRoutes() from setup_test.go
	routes := getRoutes()
	// Create a test server
	ts := httptest.NewTLSServer(routes)
	// This will only run after TestHandlder has finish execution
	defer ts.Close()

	for _, e := range theTests {
		// We only have two type of tests, for GETs and POSTs
		if e.method == "GET" {
			// Client creats a web browser or web cleint
			res, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
			}
		} else {

		}
	}

}
