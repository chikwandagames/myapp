package forms

import (
	"net/http"
	"net/url"
)

// form creates a custom form struct, embeds a url.Values object
// For holding all the info associatied with our forms
type Form struct {
	url.Values
	Error errors
}

// Initializes a for struct
// Because its a *pointer, we can access it knowing its the right form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Checks if form is not empty
// For e.g. does the form submited by the user have a field named, firstname
func (f *Form) Has(field string, r *http.Request) bool {
	// Check if the request has the field
	x := r.Form.Get(field)
	// Because its a required field, check if its empty
	if x == "" {
		return false
	}
	return true
}
