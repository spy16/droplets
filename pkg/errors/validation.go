package errors

// Common validation error type codes.
const (
	TypeMissingField = "MissingField"
	TypeInvalidValue = "InvalidValue"
)

// InvalidValue can be used to generate an error that represents an invalid
// value for the 'field'. reason should be used to add detail describing why
// the value is invalid.
func InvalidValue(field string, reason string) error {
	return WithStack(&Error{
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
		Type:    TypeMissingField,
		Message: "A required field is missing",
		Context: map[string]interface{}{
			"field": field,
		},
	})
}
