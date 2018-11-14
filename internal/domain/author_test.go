package domain_test

import (
	"fmt"
	"testing"

	"github.com/spy16/droplet/internal/domain"
	"github.com/spy16/droplet/pkg/errors"
)

func TestAuthor_Validate(suite *testing.T) {
	suite.Parallel()

	cases := []struct {
		author    domain.Author
		expectErr bool
		errType   string
	}{
		{
			author:    domain.Author{},
			expectErr: true,
			errType:   errors.TypeMissingField,
		},
		{
			author: domain.Author{
				Meta: domain.Meta{
					Kind: "Author",
					Name: "spy16",
				},
				Email: "blah.com",
			},
			expectErr: true,
			errType:   errors.TypeInvalidValue,
		},
		{
			author: domain.Author{
				Meta: domain.Meta{
					Kind: "Author",
					Name: "spy16",
				},
				Email: "spy16 <no-mail@nomail.com>",
			},
			expectErr: false,
		},
	}

	for id, cs := range cases {
		suite.Run(fmt.Sprintf("Case#%d", id), func(t *testing.T) {
			testValidation(t, cs.author, cs.expectErr, cs.errType)
		})
	}
}
