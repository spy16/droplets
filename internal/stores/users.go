package stores

import (
	"context"
	"time"

	"github.com/spy16/droplet/pkg/errors"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spy16/droplet/internal/domain"
)

// NewUsers initializes a users store with the given db handle.
func NewUsers(db *mongo.Database) *Users {
	return &Users{
		db: db,
	}
}

// Users implements UserStore interface.
type Users struct {
	db *mongo.Database
}

// Exists checks if the user identified by the given username already
// exists. Will return false in case of any error.
func (users *Users) Exists(ctx context.Context, name string) bool {
	count, err := users.db.Collection("users").Count(ctx, map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return false
	}
	return count > 0
}

// Save validates and persists the user.
func (users *Users) Save(ctx context.Context, user domain.User) (*domain.User, error) {
	user.SetDefaults()
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if _, err := users.db.Collection("users").InsertOne(ctx, user); err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByName finds a user by name. If not found, returns ResourceNotFound error.
func (users *Users) FindByName(ctx context.Context, name string) (*domain.User, error) {
	result := users.db.Collection("users").FindOne(ctx, map[string]interface{}{"name": name})
	if result == nil {
		return nil, errors.ResourceNotFound("User", name)
	}

	user := domain.User{}
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	user.SetDefaults()
	return &user, nil
}
