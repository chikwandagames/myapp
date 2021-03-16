package render

import (
	"net/http"
	"testing"

	"github.com/chikwandagames/myapp/internal/models"
)

func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// Now to test
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}

}

// Add a session to the request
func getSession() (*http.Request, error) {
	// Create a dumy request
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		return nil, err
	}

	// For writing and reading from sessions
	ctx := r.Context()
	// Put session data in the context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// Now put context back into the request
	r = r.WithContext(ctx)

	return r, nil
}
