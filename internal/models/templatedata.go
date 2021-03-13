package models

import "github.com/chikwandagames/myapp/internal/forms"

// This is made available to every page on the site
// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	// These are types that might be sent to the templates
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// Because we might send types the we don't know yet, e.g structs
	Data map[string]interface{}

	// each time a page is visited
	CSRFToken string
	// For messages
	Flash   string
	Warning string
	Error   string
	Form    *forms.Form
}
