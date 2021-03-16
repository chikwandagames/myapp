package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/chikwandagames/myapp/internal/config"
	"github.com/chikwandagames/myapp/internal/models"
)

// We need to build a request

var session *scs.SessionManager
var testApp config.AppConfig

//This function runs before any of the test are run and just before
// it closes it runs the tests
func TestMain(m *testing.M) {
	// Tell the App what kind Custom types that we want to store in a session
	// so that we can store them in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	testApp.InProduction = false

	// Initialize session
	session = scs.New()
	// How long to I want my sessions to live for, 24 hours
	session.Lifetime = 24 * time.Hour
	// Because we are storing our sessions in cookies,
	// Should the cookies persist after the browser window is closed??
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// For cookie incryption, https
	session.Cookie.Secure = false

	// For session accessibily globally
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
