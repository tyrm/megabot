package db

type Error error

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
