package domain

import (
	"fmt"
	"strings"

	"github.com/spy16/droplets/pkg/errors"
)

// Common content types.
const (
	ContentLibrary = "library"
	ContentLink    = "link"
	ContentVideo   = "video"
)

var validTypes = []string{ContentLibrary, ContentLink, ContentVideo}

// Post represents an article, link, video etc.
type Post struct {
	Meta `json:",inline" bson:",inline"`

	// Type should state the type of the content. (e.g., library,
	// video, link etc.)
	Type string `json:"type" bson:"type"`

	// Body should contain the actual content according to the Type
	// specified. (e.g. github.com/spy16/parens when Type=link)
	Body string `json:"body" bson:"body"`

	// Owner represents the name of the user who created the post.
	Owner string `json:"owner" bson:"owner"`
}

// Validate performs validation of the post.
func (post Post) Validate() error {
	if err := post.Meta.Validate(); err != nil {
		return err
	}

	if len(strings.TrimSpace(post.Body)) == 0 {
		return errors.MissingField("Body")
	}

	if len(strings.TrimSpace(post.Owner)) == 0 {
		return errors.MissingField("Owner")
	}

	if !contains(post.Type, validTypes) {
		return errors.InvalidValue("Type", fmt.Sprintf("type must be one of: %s", strings.Join(validTypes, ",")))
	}

	return nil
}

func contains(val string, vals []string) bool {
	for _, item := range vals {
		if val == item {
			return true
		}
	}
	return false
}
