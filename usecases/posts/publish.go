package posts

import (
	"context"
	"fmt"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/logger"
)

// NewPublication initializes the publication usecase.
func NewPublication(lg logger.Logger, store Store, verifier userVerifier) *Publication {
	return &Publication{
		Logger:   lg,
		store:    store,
		verifier: verifier,
	}
}

// Publication implements the publishing usecases.
type Publication struct {
	logger.Logger

	store    Store
	verifier userVerifier
}

// Publish validates and persists the post into the store.
func (pub *Publication) Publish(ctx context.Context, post domain.Post) (*domain.Post, error) {
	if err := post.Validate(); err != nil {
		return nil, err
	}

	if !pub.verifier.Exists(ctx, post.Owner) {
		return nil, errors.Unauthorized(fmt.Sprintf("user '%s' not found", post.Owner))
	}

	if pub.store.Exists(ctx, post.Name) {
		return nil, errors.Conflict("Post", post.Name)
	}

	saved, err := pub.store.Save(ctx, post)
	if err != nil {
		pub.Warnf("failed to save post to the store: %+v", err)
	}

	return saved, nil
}

// Delete removes the post from the store.
func (pub *Publication) Delete(ctx context.Context, name string) (*domain.Post, error) {
	return pub.store.Delete(ctx, name)
}
