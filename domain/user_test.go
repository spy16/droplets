package domain_test

import (
	"fmt"
	"testing"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/errors"
)

func TestUser_CheckSecret(t *testing.T) {
	password := "hello@world!"

	user := domain.User{}
	user.Secret = password
	err := user.HashSecret()
	if err != nil {
		t.Errorf("was not expecting error, got '%s'", err)
	}

	if !user.CheckSecret(password) {
		t.Errorf("CheckSecret expected to return true, but got false")
	}
}

func TestUser_Validate(suite *testing.T) {
	suite.Parallel()

	cases := []struct {
		user      domain.User
		expectErr bool
		errType   string
	}{
		{
			user:      domain.User{},
			expectErr: true,
			errType:   errors.TypeMissingField,
		},
		{
			user: domain.User{
				Meta: domain.Meta{
					Name: "spy16",
				},
				Email: "blah.com",
			},
			expectErr: true,
			errType:   errors.TypeInvalidValue,
		},
		{
			user: domain.User{
				Meta: domain.Meta{
					Name: "spy16",
				},
				Email: "spy16 <no-mail@nomail.com>",
			},
			expectErr: false,
		},
	}

	for id, cs := range cases {
		suite.Run(fmt.Sprintf("Case#%d", id), func(t *testing.T) {
			testValidation(t, cs.user, cs.expectErr, cs.errType)
		})
	}
}
