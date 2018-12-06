package mongo

import (
	"context"
	"time"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const colPosts = "posts"

// NewPostStore initializes the Posts store with given mongo db handle.
func NewPostStore(db *mgo.Database) *PostStore {
	return &PostStore{
		db: db,
	}
}

// PostStore manages persistence and retrieval of posts.
type PostStore struct {
	db *mgo.Database
}

// Exists checks if a post exists by name.
func (posts *PostStore) Exists(ctx context.Context, name string) bool {
	col := posts.db.C(colPosts)

	count, err := col.Find(bson.M{"name": name}).Count()
	if err != nil {
		return false
	}
	return count > 0
}

// Get finds a post by name.
func (posts *PostStore) Get(ctx context.Context, name string) (*domain.Post, error) {
	col := posts.db.C(colPosts)

	post := domain.Post{}
	if err := col.Find(bson.M{"name": name}).One(&post); err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.ResourceNotFound("Post", name)
		}
		return nil, errors.Wrapf(err, "failed to fetch post")
	}

	post.SetDefaults()
	return &post, nil
}

// Save validates and persists the post.
func (posts *PostStore) Save(ctx context.Context, post domain.Post) (*domain.Post, error) {
	post.SetDefaults()
	if err := post.Validate(); err != nil {
		return nil, err
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	col := posts.db.C(colPosts)
	if err := col.Insert(post); err != nil {
		return nil, err
	}
	return &post, nil
}

// Delete removes one post identified by the name.
func (posts *PostStore) Delete(ctx context.Context, name string) (*domain.Post, error) {
	col := posts.db.C(colPosts)

	ch := mgo.Change{
		Remove:    true,
		ReturnNew: true,
		Upsert:    false,
	}
	post := domain.Post{}
	_, err := col.Find(bson.M{"name": name}).Apply(ch, &post)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.ResourceNotFound("Post", name)
		}
		return nil, err
	}

	return &post, nil
}
