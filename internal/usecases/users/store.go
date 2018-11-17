package users

import (
	"context"

	"github.com/spy16/droplet/internal/domain"
)

// UserStore implementation is responsible for managing persistence of
// users.
type UserStore interface {
	Exists(ctx context.Context, name string) bool
	Save(ctx context.Context, user domain.User) (*domain.User, error)
	FindByName(ctx context.Context, name string) (*domain.User, error)
}
