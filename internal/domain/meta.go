package domain

import (
	"strings"
	"time"

	"github.com/spy16/droplet/pkg/errors"
)

// Meta represents metadata about different entities.
type Meta struct {
	// Kind represents the type of the object.
	Kind string `json:"kind"`

	// Name represents a unique name/identifier for the object.
	Name string `json:"name"`

	// Labels can contain additional metadata about the object.
	Labels map[string]string `json:"labels,omitempty"`

	// CreateAt represents the time at which this object was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// UpdatedAt represents the time at which this object was last
	// modified.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// SetDefaults sets sensible defaults on meta.
func (meta *Meta) SetDefaults() {
	meta.Labels = map[string]string{}
	if meta.CreatedAt.IsZero() {
		meta.CreatedAt = time.Now()
		meta.UpdatedAt = time.Now()
	}
}

// Validate performs basic validation of the metadata.
func (meta Meta) Validate() error {
	switch {
	case empty(meta.Kind):
		return errors.MissingField("Kind")
	case empty(meta.Name):
		return errors.MissingField("Name")
	}
	return nil
}

func empty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
