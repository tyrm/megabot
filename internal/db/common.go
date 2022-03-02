package db

import "context"

// Common wraps common database functions
type Common interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Create stores the object
	Create(context.Context, interface{}) Error
}
