package chatbot

import "fmt"

// Error is a chatbot specific error
type Error error

var (
	// ErrAlreadyRunning is returned when a service worker is asked start, but it's not in a stopped state.
	ErrAlreadyRunning Error = fmt.Errorf("service is already running")
	// ErrAPIError is returned when calling the service's api returns an error
	ErrAPIError Error = fmt.Errorf("api error")
	// ErrIDExists is returned when a service worker is asked to be added, but it already exists.
	ErrIDExists Error = fmt.Errorf("service id already exists")
)
