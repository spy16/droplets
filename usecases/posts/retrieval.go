package posts

import (
	"context"
	"errors"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/logger"
)

// NewRetriever initializes the retrieval usecase with given store.
func NewRetriever(lg logger.Logger, store Store) *Retriever {
	return &Retriever{
		Logger: lg,
		store:  store,
	}
}

// Retriever provides retrieval related usecases.
type Retriever struct {
	logger.Logger

	store Store
}

// Get finds a post by its name.
func (ret *Retriever) Get(ctx context.Context, name string) (*domain.Post, error) {
	return ret.store.Get(ctx, name)
}

// Search finds all the posts matching the parameters in the query.
func (ret *Retriever) Search(ctx context.Context, query Query) ([]domain.Post, error) {
	return nil, errors.New("not implemented")
}

// Query represents parameters for executing a search. Zero valued fields
// in the query will be ignored.
type Query struct {
	Name  string   `json:"name,omitempty"`
	Owner string   `json:"owner,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}
