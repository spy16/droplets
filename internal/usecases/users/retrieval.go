package users

import (
	"context"

	"github.com/spy16/droplet/internal/domain"
	"github.com/spy16/droplet/pkg/logger"
)

// NewRetriever initializes an instance of Retriever with given store.
func NewRetriever(lg logger.Logger, store UserStore) *Retriever {
	return &Retriever{
		Logger: lg,
		store:  store,
	}
}

// Retriever provides functions for retrieving user and user info.
type Retriever struct {
	logger.Logger

	store UserStore
}

// Get finds a user by name.
func (ret *Retriever) Get(ctx context.Context, name string) (*domain.User, error) {
	user, err := ret.store.FindByName(ctx, name)
	if err != nil {
		ret.Debugf("failed to find user with name '%s': %v", name, err)
		return nil, err
	}

	return user, nil
}
