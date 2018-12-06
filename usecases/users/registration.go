package users

import (
	"context"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/logger"
)

// NewRegistrar initializes a Registration service object.
func NewRegistrar(lg logger.Logger, store Store) *Registrar {
	return &Registrar{
		Logger: lg,
		store:  store,
	}
}

// Registrar provides functions for user registration.
type Registrar struct {
	logger.Logger

	store Store
}

// Register creates a new user in the system using the given user object.
func (reg *Registrar) Register(ctx context.Context, user domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if len(user.Secret) < 8 {
		return nil, errors.InvalidValue("Secret", "secret must have 8 or more characters")
	}

	if reg.store.Exists(ctx, user.Name) {
		return nil, errors.Conflict("User", user.Name)
	}

	if err := user.HashSecret(); err != nil {
		return nil, err
	}

	saved, err := reg.store.Save(ctx, user)
	if err != nil {
		reg.Logger.Warnf("failed to save user object: %v", err)
		return nil, err
	}

	saved.Secret = ""
	return saved, nil
}
