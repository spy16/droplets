package posts

import (
	"context"

	"github.com/spy16/droplets/internal/domain"
)

// PostStore implementation is responsible for managing persistance of posts.
type PostStore interface {
	Get(ctx context.Context, name string) (*domain.Post, error)
	Exists(ctx context.Context, name string) bool
	Save(ctx context.Context, post domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, name string) (*domain.Post, error)
}

// userVerifier is responsible for verifying existence of a user.
type userVerifier interface {
	Exists(ctx context.Context, name string) bool
}
