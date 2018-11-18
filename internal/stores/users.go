package stores

import (
	"context"
	"time"

	"github.com/spy16/droplet/internal/domain"
	"github.com/spy16/droplet/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewUsers initializes a users store with the given db handle.
func NewUsers(db *mgo.Database) *Users {
	return &Users{
		db: db,
	}
}

// Users implements UserStore interface.
type Users struct {
	db *mgo.Database
}

// Exists checks if the user identified by the given username already
// exists. Will return false in case of any error.
func (users *Users) Exists(ctx context.Context, name string) bool {
	col := users.db.C("users")

	count, err := col.Find(bson.M{"name": name}).Count()
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

	col := users.db.C("users")
	if err := col.Insert(user); err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByName finds a user by name. If not found, returns ResourceNotFound error.
func (users *Users) FindByName(ctx context.Context, name string) (*domain.User, error) {
	col := users.db.C("users")

	user := domain.User{}
	if err := col.Find(bson.M{"name": name}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.ResourceNotFound("User", name)
		}
		return nil, errors.Wrapf(err, "failed to fetch user")
	}

	user.SetDefaults()
	return &user, nil
}

// FindAll finds all users matching the tags.
func (users *Users) FindAll(ctx context.Context, tags []string, limit int) ([]domain.User, error) {
	col := users.db.C("users")

	filter := bson.M{}
	if len(tags) > 0 {
		filter["tags"] = bson.M{
			"$in": tags,
		}
	}

	matches := []domain.User{}
	if err := col.Find(filter).Limit(limit).All(&matches); err != nil {
		return nil, errors.Wrapf(err, "failed to query for users")
	}
	return matches, nil
}
