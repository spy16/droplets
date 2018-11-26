package domain

import (
	"net/mail"

	"github.com/spy16/droplets/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User represents information about registered users.
type User struct {
	Meta `json:",inline,omitempty" bson:",inline"`

	// Email should contain a valid email of the user.
	Email string `json:"email,omitempty" bson:"email"`

	// Secret represents the user secret.
	Secret string `json:"secret,omitempty" bson:"secret"`
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

// HashSecret creates bcrypt hash of the password.
func (user *User) HashSecret() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Secret), 4)
	if err != nil {
		return err
	}
	user.Secret = string(bytes)
	return nil
}

// CheckSecret compares the cleartext password with the hash.
func (user User) CheckSecret(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Secret), []byte(password))
	return err == nil
}
