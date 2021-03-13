package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is a custom middleware that prints to the
// console each time a page is accessed
func WriteToConsole(next http.Handler) http.Handler {
	// The anolymous func is cast to a http.Handler func
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
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
