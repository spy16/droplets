package errors

// Common authorization related errors
const (
	TypeUnauthorized = "Unauthorized"
)

// Unauthorized can be used to generate an error that represents an unauthorized
// request.
func Unauthorized(action string) error {
	return WithStack(&Error{
		Type:    TypeUnauthorized,
		Message: "You are not authorized to perform the requested action",
		Context: map[string]interface{}{
			"action": action,
		},
	})
}
