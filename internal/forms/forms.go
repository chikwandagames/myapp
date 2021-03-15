package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// form creates a custom form struct, embeds a url.Values object
// For holding all the info associatied with our forms
type Form struct {
	url.Values
	Errors errors
}

// Initializes a for struct
// Because its a *pointer, we can access it knowing its the right form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Checks for required fields
func (f *Form) Required(fields ...string) {
	// The value here will be e.g. firstname, lastname, email, phone ...
	for _, field := range fields {
		val := f.Get(field)
		// Optional field should handled here
		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This field is requird")
		}
	}
}

// Checks if form is not empty
// For e.g. does the form submited by the user have a field named, firstname
func (f *Form) Has(field string, r *http.Request) bool {
	// Check if the request has the field
	x := r.Form.Get(field)
	// Because its a required field, check if its empty
	if x != "" {
		return true
	}
	// f.Errors.Add(field, "This field is required")
	return false
}

// Valid returns true if there are no errors
func (f *Form) Valid() bool {
	// Return true, if no errors
	return len(f.Errors) == 0
}

// MinLength checks for string minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}
