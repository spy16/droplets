package errors

import "net/http"

// Common authorization related errors
const (
	TypeUnauthorized = "Unauthorized"
)

// Unauthorized can be used to generate an error that represents an unauthorized
// request.
func Unauthorized(reason string) error {
	return WithStack(&Error{
		Code:    http.StatusUnauthorized,
		Type:    TypeUnauthorized,
		Message: "You are not authorized to perform the requested action",
		Context: map[string]interface{}{
			"reason": reason,
		},
	})
}
