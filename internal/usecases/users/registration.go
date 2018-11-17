package users

import (
	"context"

	"github.com/spy16/droplet/internal/domain"
	"github.com/spy16/droplet/pkg/errors"
)

// NewRegistration initializes a Registration service object.
func NewRegistration(store UserStore) *Registration {
	return &Registration{
		store: store,
	}
}

// Registration provides functions for user registration.
type Registration struct {
	store UserStore
}

// Register creates a new user in the system using the given user object.
func (reg *Registration) Register(ctx context.Context, user domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if reg.store.Exists(ctx, user.Name) {
		return nil, errors.Conflict("User", user.Name)
	}

	return reg.store.Save(ctx, user)
}
