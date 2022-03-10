package webapp

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// LogoutGetHandler logs a user out
func (m *Module) LogoutGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init Session
	us := r.Context().Value(sessionKey).(*sessions.Session)

	// Set user to nil
	us.Values["user"] = nil
	err := us.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, pathBase+pathLogin, http.StatusFound)
}
