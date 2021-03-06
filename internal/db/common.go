package db

import "context"

// Common wraps common database functions
type Common interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Create stores the object
	Create(ctx context.Context, i any) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
	// LoadTestData adds test data to the database
	LoadTestData(ctx context.Context) Error
}
