package webapp

import "net/http"

// HomeTemplate contains the variables for the "home" template.
type HomeTemplate struct {
	templateCommon
}

// HomeGetHandler serves the home page
func (m *Module) HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init template variables
	tmplVars := &HomeTemplate{
		templateCommon{
			PageTitle: "Home",
		},
	}

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.executeTemplate(w, "home", tmplVars)
	if err != nil {
		logger.Errorf("could not render home template: %s", err.Error())
	}

}
