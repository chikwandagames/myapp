package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/justinas/nosurf"

	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/models"
	"github.com/chikwandagames/myapp/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// This file setup_test.go will run before the tests run

var app config.AppConfig
var session *scs.SessionManager

// A FuncMap is a map of function that can be used in a template
var functions = template.FuncMap{}

// Because we do not running our tests from the root of the app
// we need to create a path
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {
	// Tell the App what kind Custom types that we want to store in a session
	// so that we can store them in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	// Initialize session
	session = scs.New()
	// How long to I want my sessions to live for, 24 hours
	session.Lifetime = 24 * time.Hour
	// Because we are storing our sessions in cookies,
	// Should the cookies persist after the browser window is closed??
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// For cookie incryption, https
	session.Cookie.Secure = app.InProduction

	// For session accessibily globally
	app.Session = session

	// Load Templates cache, when the app starts
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// Store template cache in the app
	app.TemplateCache = tc
	// We use this to load Templates from disc or cache, depending
	// on weather we in development of production mode
	app.UseCache = true

	repo := NewRepo(&app)
	// Pass ther repo to the handlers
	NewHandlers(repo)

	// Pass the template cache to render package
	// giving the render package access to app (template)
	render.NewTemplates(&app)

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

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/make-reservation", Repo.Reservation)
	mux.Get("/generals", Repo.Generals)
	mux.Get("/majors", Repo.Majors)
	mux.Get("/search-availability", Repo.Availability)
	mux.Get("/reservation-summary", Repo.ReservationSummary)
	mux.Get("/contact", Repo.Contact)

	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	mux.Post("/make-reservation", Repo.PostReservation)

	// For access to static files, image...
	fileServer := http.FileServer(http.Dir("./static/"))
	// Strip out /static and use fileServer
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

// NoSurf adds CSRF protection to all POST requests
// NoSurf is for forms, security tocken, cross site request forgery token
// Its a hidden field in a form, string of random numbrer, which change
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// It uses cookies to make sure, the token it generates a cookie
	// is available on a per page basis
	// So we create a new cookie
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		// "/" means the entire site
		Path:     "/",
		Secure:   app.InProduction, // Not https
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
// LoadAndSave provides middleware which automatically loads
// and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTemplateCache creates template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {

	// myCache is template cache to hold all templates, and is created
	// at the start of the application, key string, value template pointer
	myCache := map[string]*template.Template{}

	// Add files with .page.html suffix to pages var
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// Because we don't want the full path, we remove the ./templates/
		name := filepath.Base(page)
		// fmt.Println("Page is currently", page)
		// fmt.Println("name is currently", name)
		// Create a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		// If we find a match, i.e. a file with the suffix, ./templates/*.layout.html
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		// Now add the template to the cache
		myCache[name] = ts

	}
	return myCache, nil
}
