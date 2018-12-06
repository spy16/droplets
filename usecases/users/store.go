package users

import (
	"context"

	"github.com/spy16/droplets/domain"
)

// Store implementation is responsible for managing persistence of
// users.
type Store interface {
	Exists(ctx context.Context, name string) bool
	Save(ctx context.Context, user domain.User) (*domain.User, error)
	FindByName(ctx context.Context, name string) (*domain.User, error)
	FindAll(ctx context.Context, tags []string, limit int) ([]domain.User, error)
}
