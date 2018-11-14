package domain

import (
	"net/mail"

	"github.com/spy16/droplet/pkg/errors"
)

// Author represents information about registered authors.
type Author struct {
	Meta `json:",inline"`

	// Email should contain a valid email of the author.
	Email string
}

// Validate performs basic validation of author information.
func (author Author) Validate() error {
	if err := author.Meta.Validate(); err != nil {
		return err
	}

	_, err := mail.ParseAddress(author.Email)
	if err != nil {
		return errors.InvalidValue("Email", err.Error())
	}

	return nil
}
