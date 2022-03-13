package web

// Module represents a module that can be added to a web server
type Module interface {
	Name() string
	Route(s Server) error
}
