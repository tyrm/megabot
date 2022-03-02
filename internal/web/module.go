package web

type Module interface {
	Route(s Server) error
}
