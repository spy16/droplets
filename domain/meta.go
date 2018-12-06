package domain

import (
	"strings"
	"time"

	"github.com/spy16/droplets/pkg/errors"
)

// Meta represents metadata about different entities.
type Meta struct {
	// Name represents a unique name/identifier for the object.
	Name string `json:"name" bson:"name"`

	// Tags can contain additional metadata about the object.
	Tags []string `json:"tags,omitempty" bson:"tags"`

	// CreateAt represents the time at which this object was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// UpdatedAt represents the time at which this object was last
	// modified.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// SetDefaults sets sensible defaults on meta.
func (meta *Meta) SetDefaults() {
	if meta.CreatedAt.IsZero() {
		meta.CreatedAt = time.Now()
		meta.UpdatedAt = time.Now()
	}
}

// Validate performs basic validation of the metadata.
func (meta Meta) Validate() error {
	switch {
	case empty(meta.Name):
		return errors.MissingField("Name")
	}
	return nil
}

func empty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
