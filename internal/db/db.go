package db

// DB represents a database client
type DB interface {
	Common

	Chatbot
	User
}
