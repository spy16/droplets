package users

import (
	"context"

	"github.com/spy16/droplet/internal/domain"
)

// Retriever provides functions for retrieving user and user info.
type Retriever struct {
	store UserStore
}

// GetUser finds a user by name.
func (ret *Retriever) GetUser(ctx context.Context, name string) (*domain.User, error) {
	return ret.store.FindByName(ctx, name)
}
