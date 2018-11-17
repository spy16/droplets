package domain

import (
	"net/mail"

	"github.com/spy16/droplet/pkg/errors"
)

// User represents information about registered users.
type User struct {
	Meta `json:",inline,omitempty" bson:",inline"`

	// Email should contain a valid email of the user.
	Email string `json:"email,omitempty" bson:"email"`
}

// Validate performs basic validation of user information.
func (user User) Validate() error {
	if err := user.Meta.Validate(); err != nil {
		return err
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.InvalidValue("Email", err.Error())
	}

	return nil
}
