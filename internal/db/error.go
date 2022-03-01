package db

import "fmt"

type Error error

var (
	// ErrNoEntries is returned when a caller expected an entry for a query, but none was found.
	ErrNoEntries Error = fmt.Errorf("no entries")
	// ErrMultipleEntries is returned when a caller expected ONE entry for a query, but multiples were found.
	ErrMultipleEntries Error = fmt.Errorf("multiple entries")
	// ErrUnknown denotes an unknown database error.
	ErrUnknown Error = fmt.Errorf("unknown error")
)

// ErrAlreadyExists is returned when a caller tries to insert a database entry that already exists in the db.
type ErrAlreadyExists struct {
	message string
}

func (e *ErrAlreadyExists) Error() string {
	return e.message
}

func NewErrAlreadyExists(msg string) error {
	return &ErrAlreadyExists{message: msg}
}
