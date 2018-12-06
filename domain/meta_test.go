package domain_test

import (
	"fmt"
	"testing"

	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/errors"
)

func TestMeta_Validate(suite *testing.T) {
	suite.Parallel()

	cases := []struct {
		meta      domain.Meta
		expectErr bool
		errType   string
	}{}

	for id, cs := range cases {
		suite.Run(fmt.Sprintf("Case#%d", id), func(t *testing.T) {
			testValidation(t, cs.meta, cs.expectErr, cs.errType)
		})
	}

}

func testValidation(t *testing.T, validator validatable, expectErr bool, errType string) {
	err := validator.Validate()
	if err != nil {
		if !expectErr {
			t.Errorf("unexpected error: %s", err)
			return
		}

		if actualType := errors.Type(err); actualType != errType {
			t.Errorf("expecting error type '%s', got '%s'", errType, actualType)
		}
		return
	}

	if expectErr {
		t.Errorf("was expecting an error of type '%s', got nil", errType)
		return
	}
}

type validatable interface {
	Validate() error
}
