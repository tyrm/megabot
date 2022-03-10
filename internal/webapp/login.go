package webapp

import "net/http"

// LoginTemplate contains the variables for the "login" template.
type LoginTemplate struct {
	templateCommon
}

// LoginGetHandler serves the login page
func (m *Module) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init template variables
	tmplVars := &LoginTemplate{
		templateCommon{
			PageTitle: "Login",
		},
	}

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.executeTemplate(w, "login", tmplVars)
	if err != nil {
		logger.Errorf("could not render login template: %s", err.Error())
	}
}
