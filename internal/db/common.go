package db

import "context"

// Common wraps common database functions
type Common interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Put stores the object
	Put(context.Context, interface{}) Error
}
