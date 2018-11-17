package errors

import "net/http"

// Common validation error type codes.
const (
	TypeInvalidRequest = "InvalidRequest"
	TypeMissingField   = "MissingField"
	TypeInvalidValue   = "InvalidValue"
)

// Validation returns an error that can be used to represent an invalid request.
func Validation(reason string) error {
	return WithStack(&Error{
		Code:    http.StatusBadRequest,
		Type:    TypeInvalidRequest,
		Message: reason,
		Context: map[string]interface{}{},
	})
}

// InvalidValue can be used to generate an error that represents an invalid
// value for the 'field'. reason should be used to add detail describing why
// the value is invalid.
func InvalidValue(field string, reason string) error {
	return WithStack(&Error{
		Code:    http.StatusBadRequest,
		Type:    TypeInvalidValue,
		Message: "A parameter has invalid value",
		Context: map[string]interface{}{
			"field":  field,
			"reason": reason,
		},
	})
}

// MissingField can be used to generate an error that represents
// a empty value for a required field.
func MissingField(field string) error {
	return WithStack(&Error{
		Code:    http.StatusBadRequest,
		Type:    TypeMissingField,
		Message: "A required field is missing",
		Context: map[string]interface{}{
			"field": field,
		},
	})
}
