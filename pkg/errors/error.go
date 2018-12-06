package errors

import (
	"fmt"
	"io"
)

// Error is a generic error representation with some fields to provide additional
// context around the error.
type Error struct {
	// Code can represent an http error code.
	Code int `json:"-"`

	// Type should be an error code to identify the error. Type and Context together
	// should provide enough context for robust error handling on client side.
	Type string `json:"type,omitempty"`

	// Context can contain additional information describing the error. Context will
	// be exposed only in API endpoints so that clients be integrated effectively.
	Context map[string]interface{} `json:"context,omitempty"`

	// Message should be a user-friendly error message which can be shown to the
	// end user without further modifications. However, clients are free to modify
	// this (e.g., for enabling localization), or augment this message with the
	// information available in the context before rendering a message to the end
	// user.
	Message string `json:"message,omitempty"`

	// original can contain an underlying error if any. This value will be returned
	// by the Cause() method.
	original error

	// stack will contain a minimal stack trace which can be used for logging and
	// debugging. stack should not be examined to handle errors.
	stack stack
}

// Cause returns the underlying error if any.
func (err Error) Cause() error {
	return err.original
}

func (err Error) Error() string {
	if origin := err.Cause(); origin != nil {
		return fmt.Sprintf("%s: %s: %s", origin, err.Type, err.Message)
	}

	return fmt.Sprintf("%s: %s", err.Type, err.Message)
}

// Format implements fmt.Formatter interface.
func (err Error) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, err.Error())
			err.stack.Format(st, verb)
		} else {
			fmt.Fprintf(st, "%s: ", err.Type)
			for key, val := range err.Context {
				fmt.Fprintf(st, "%s='%s' ", key, val)
			}
		}
	case 's':
		io.WriteString(st, err.Error())
	case 'q':
		fmt.Fprintf(st, "%q", err.Error())
	}
}
