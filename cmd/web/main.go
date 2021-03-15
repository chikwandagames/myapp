package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/handlers"
	"github.com/chikwandagames/myapp/internal/models"
	"github.com/chikwandagames/myapp/internal/render"
)

const portNumber = ":8080"

// Make variable visible throughout the main package
var app config.AppConfig
var session *scs.SessionManager

// if A function or variable begins with a CAPITAL, this means its accessible outside
// that package

func main() {
	// Tell the App what kind Custom types that we want to store in a session
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
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// Store template cache in the app
	app.TemplateCache = tc
	// We use this to load Templates from disc or cache, depending
	// on weather we in development of production mode
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	// Pass ther repo to the handlers
	handlers.NewHandlers(repo)

	// Pass the template cache to render package
	// giving the render package access to app (template)
	render.NewTemplates(&app)

	fmt.Printf("Staring application on port %s \n", portNumber)

	// Start a webserver

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	// If theres an error
	log.Fatal(err)

}
