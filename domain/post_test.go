package domain_test

import (
	"fmt"
	"testing"

	"github.com/spy16/droplets/domain"
)

func TestPost_Validate(suite *testing.T) {
	suite.Parallel()

	validMeta := domain.Meta{
		Name: "hello",
	}

	cases := []struct {
		post      domain.Post
		expectErr bool
	}{
		{
			post:      domain.Post{},
			expectErr: true,
		},
		{
			post: domain.Post{
				Meta: validMeta,
			},
			expectErr: true,
		},
		{
			post: domain.Post{
				Meta: validMeta,
				Body: "hello world post!",
			},
			expectErr: true,
		},
		{
			post: domain.Post{
				Meta:  validMeta,
				Type:  "blah",
				Owner: "spy16",
				Body:  "hello world post!",
			},
			expectErr: true,
		},
		{
			post: domain.Post{
				Meta: validMeta,
				Type: domain.ContentLibrary,
				Body: "hello world post!",
			},
			expectErr: true,
		},
		{
			post: domain.Post{
				Meta:  validMeta,
				Type:  domain.ContentLibrary,
				Body:  "hello world post!",
				Owner: "spy16",
			},
			expectErr: false,
		},
	}

	for id, cs := range cases {
		suite.Run(fmt.Sprintf("#%d", id), func(t *testing.T) {
			err := cs.post.Validate()
			if err != nil {
				if !cs.expectErr {
					t.Errorf("was not expecting error, got '%s'", err)
				}
				return
			}

			if cs.expectErr {
				t.Errorf("was expecting error, got nil")
			}
		})
	}
}
