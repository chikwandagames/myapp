package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/forms"
	"github.com/chikwandagames/myapp/internal/models"
	"github.com/chikwandagames/myapp/internal/render"
)

// Repo is the repository used by handlers
// Here we use a repository pattern, to help us swap components
// in and out off the application with minimal changes to the code.
// We want the handlers to have access to config
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page, Handler function
// In order for a function to respond to requests from a web browser
// it has to handle 2 params, ResponseWriter and Request
// Home is ...
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Get user IP and store in session
	remoteIP := r.RemoteAddr
	// Takes a context, key and value of item to store in session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})

}

// About is the handler for the about page
// Adding the receiver, gives About handler access to everything inside
// the repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("about handler called")
	// render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there"

	// Get value stored in session
	// Use GetString(), that way no need for casting
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation renders the make reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// To repopulate input field if a user enters the wrong info, so that the
	// user does not need to retype every word
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	// Key should match key in Post Reservation
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		// New(nil) is and empty form
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// This data might contain bad data
	reservation := models.Reservation{
		// Get data submitted from the form
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	// var form now hold all data from the form as a pointer
	form := forms.New(r.PostForm)

	// Now validate, does form have ...
	// form.Has("first_name", r)
	// Pass the required fields, for validation
	form.Required("first_name", "last_name", "email")
	form.MinLength("phone", 5, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		// Rerender the form
		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		// If not valid return
		return
	}

	// Redirect to /reservation-summary, and use session to forward forward form data
	// to Reservation-summary page
	// Put(), takes a context, key and value of item to store in session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// Generals renders room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// Get data posted the from form, Get("start") the string should match the name
	// of the input tag in the html, by defauld when you get data out of a form
	// it is a string
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	w.Write([]byte("start date: " + startDate + ", end date: " + endDate))
}

// For a one off struct, by convention put as close as possible to the code
// that uses it
type jsonResponse struct {
	// To allow for export the member names have to start with a capital
	// `json:"ok"` member name in JSON
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends json response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// Generate json
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	// Convert to JSON and send back, indent by 5 spaces
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	// log.Println(string(out))
	// Header that tell the browser what kind of response we are sending
	w.Header().Set("Content-type", "application/json")
	w.Write(out)

}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	// Get data received from the redirect from make-reservation form.
	// .(models.Reservation) asserts that the data is of type models.Reservaion
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("Cannot get item from session")
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})

}
