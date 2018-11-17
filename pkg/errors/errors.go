package errors

import (
	"fmt"
	"net/http"
)

// TypeUnknown represents unknown error type.
const TypeUnknown = "Unknown"

// New returns an error object with formatted error message generated using
// the arguments.
func New(msg string, args ...interface{}) error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Type:    TypeUnknown,
		Message: fmt.Sprintf(msg, args...),
		stack:   callStack(3),
	}
}

// Type attempts converting the err to Error type and extracts error Type.
// If conversion not possible, returns TypeUnknown.
func Type(err error) string {
	if e, ok := err.(*Error); ok {
		return e.Type
	}
	return TypeUnknown
}

// Wrapf wraps the given err with formatted message and returns a new error.
func Wrapf(err error, msg string, args ...interface{}) error {
	return WithStack(&Error{
		Code:     http.StatusInternalServerError,
		Type:     TypeUnknown,
		Message:  fmt.Sprintf(msg, args...),
		original: err,
	})
}

// WithStack annotates the given error with stack trace and returns the wrapped
// error.
func WithStack(err error) error {
	var wrappedErr Error
	if e, ok := err.(*Error); ok {
		wrappedErr = *e
	} else {
		wrappedErr.Type = TypeUnknown
		wrappedErr.Message = "Something went wrong"
		wrappedErr.original = err
	}

	wrappedErr.stack = callStack(3)
	return &wrappedErr
}

// Cause returns the underlying error if the given error is wrapping another error.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	return err
}
