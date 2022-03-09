package webapp

import "net/http"

const pathHome = "/"

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

	err := m.templates.ExecuteTemplate(w, "home", tmplVars)
	if err != nil {
		logger.Errorf("could not render home template: %s", err.Error())
	}

}
