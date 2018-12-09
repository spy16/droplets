package middlewares

import (
	"context"
	"net/http"

	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/logger"
	"github.com/spy16/droplets/pkg/render"
)

var authUser = ctxKey("user")

// WithBasicAuth adds Basic authentication checks to the handler. Basic Auth header
// will be extracted from the request and verified using the verifier.
func WithBasicAuth(lg logger.Logger, next http.Handler, verifier UserVerifier) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		name, secret, ok := req.BasicAuth()
		if !ok {
			render.JSON(wr, http.StatusUnauthorized, errors.Unauthorized("Basic auth header is not present"))
			return
		}

		verified := verifier.VerifySecret(req.Context(), name, secret)
		if !verified {
			wr.WriteHeader(http.StatusUnauthorized)
			render.JSON(wr, http.StatusUnauthorized, errors.Unauthorized("Invalid username or secret"))
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), authUser, name))
		next.ServeHTTP(wr, req)
	})
}

// User extracts the username injected into the context by the auth middleware.
func User(req *http.Request) (string, bool) {
	val := req.Context().Value(authUser)
	if userName, ok := val.(string); ok {
		return userName, true
	}

	return "", false
}

type ctxKey string

// UserVerifier implementation is responsible for verifying the name-secret pair.
type UserVerifier interface {
	VerifySecret(ctx context.Context, name, secret string) bool
}

// UserVerifierFunc implements UserVerifier.
type UserVerifierFunc func(ctx context.Context, name, secret string) bool

// VerifySecret delegates call to the wrapped function.
func (uvf UserVerifierFunc) VerifySecret(ctx context.Context, name, secret string) bool {
	return uvf(ctx, name, secret)
}
