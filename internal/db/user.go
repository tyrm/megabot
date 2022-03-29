package db

import (
	"context"
	"github.com/tyrm/megabot/internal/models"
)

// User contains functions related to users.
type User interface {
	// ReadUserByID returns one user.
	ReadUserByID(ctx context.Context, id int64) (*models.User, Error)
	// ReadUserByEmail returns one user.
	ReadUserByEmail(ctx context.Context, email string) (*models.User, Error)
}
