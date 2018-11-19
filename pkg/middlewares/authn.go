package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/logger"
)

// WithAuthentication adds Basic authentication checks to the handler. Basic Auth header
// will be extracted from the request and verified using the verifier.
func WithAuthentication(verifier userVerifier, lg logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		name, secret, ok := req.BasicAuth()
		if !ok {
			json.NewEncoder(wr).Encode(errors.Unauthorized("Basic auth header is not present"))
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		verified := verifier.VerifySecret(req.Context(), name, secret)
		if !verified {
			json.NewEncoder(wr).Encode(errors.Unauthorized("Invalid username or secret"))
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(wr, req)
	})
}

type userVerifier interface {
	VerifySecret(ctx context.Context, name, secret string) bool
}
