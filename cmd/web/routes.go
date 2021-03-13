package main

import (
	"net/http"

	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Here we user the bmizerany pat external package for routing

func routes(app *config.AppConfig) http.Handler {
	// Create a multiplexer, which is an http handler
	mux := chi.NewRouter()
	// Middlewere
	// Middleware lets you process a request as it comes into the application
	// and performes some action on it

	// Help application, Gracefully absorb panics and prints the stack trace
	mux.Use(middleware.Recoverer)
	// Custom middleware
	// mux.Use(WriteToConsole)
	// This middleware says, ignore any request that is a post, that doesn't have
	// a propper cross sight request forgery tocken
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	// For access to static files, image...
	fileServer := http.FileServer(http.Dir("./static/"))
	// Strip out /static and use fileServer
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
