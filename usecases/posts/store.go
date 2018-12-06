package posts

import (
	"context"

	"github.com/spy16/droplets/domain"
)

// Store implementation is responsible for managing persistance of posts.
type Store interface {
	Get(ctx context.Context, name string) (*domain.Post, error)
	Exists(ctx context.Context, name string) bool
	Save(ctx context.Context, post domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, name string) (*domain.Post, error)
}

// userVerifier is responsible for verifying existence of a user.
type userVerifier interface {
	Exists(ctx context.Context, name string) bool
}
