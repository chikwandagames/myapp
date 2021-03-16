package main

import (
	"fmt"
	"testing"

	"github.com/chikwandagames/myapp/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	// Check if v is of type *chi.Mux which is what is returned by routes function
	case *chi.Mux:
		// do nothing: test passed

	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}

}