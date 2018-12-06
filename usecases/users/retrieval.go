package users

import (
	"context"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/logger"
)

// NewRetriever initializes an instance of Retriever with given store.
func NewRetriever(lg logger.Logger, store Store) *Retriever {
	return &Retriever{
		Logger: lg,
		store:  store,
	}
}

// Retriever provides functions for retrieving user and user info.
type Retriever struct {
	logger.Logger

	store Store
}

// Search finds all users matching the tags.
func (ret *Retriever) Search(ctx context.Context, tags []string, limit int) ([]domain.User, error) {
	users, err := ret.store.FindAll(ctx, tags, limit)
	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].Secret = ""
	}

	return users, nil
}

// Get finds a user by name.
func (ret *Retriever) Get(ctx context.Context, name string) (*domain.User, error) {
	return ret.findUser(ctx, name, true)
}

// VerifySecret finds the user by name and verifies the secret against the has found
// in the store.
func (ret *Retriever) VerifySecret(ctx context.Context, name, secret string) bool {
	user, err := ret.findUser(ctx, name, false)
	if err != nil {
		return false
	}

	return user.CheckSecret(secret)
}

func (ret *Retriever) findUser(ctx context.Context, name string, stripSecret bool) (*domain.User, error) {
	user, err := ret.store.FindByName(ctx, name)
	if err != nil {
		ret.Debugf("failed to find user with name '%s': %v", name, err)
		return nil, err
	}

	if stripSecret {
		user.Secret = ""
	}
	return user, nil
}
