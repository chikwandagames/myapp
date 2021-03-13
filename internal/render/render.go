package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/models"
	"github.com/justinas/nosurf"
)

// A FuncMap is a map of function that can be used in a template
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData is for adding data that we need present on every page
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplateThree is ...
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	// Use the useCache bool to to decide, when in development mode,
	// dont use template cache, load template from disc
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Retrieve the template out of the map
	// Use comma ok idiom to check if the variable exists
	t, ok := tc[tmpl]
	if !ok {
		// exit
		log.Fatal("could not get template from cache")
	}

	// We need to read the template from disk
	// We need to put that parsed template in memory into some bytes
	// Buffer to hold bytes
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	// Execute the template and store in buf variable
	_ = t.Execute(buf, td)
	// Write to ResponseWriter
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// myCache is template cache to hold all templates, and is created
	// at the start of the application, key string, value template pointer
	myCache := map[string]*template.Template{}

	// Add files with .page.html suffix to pages var
	pages, err := filepath.Glob("./templates/*.page.html")
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

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// If we find a match, i.e. a file with the suffix, ./templates/*.layout.html
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		// Now add the template to the cache
		myCache[name] = ts

	}
	return myCache, nil
}
