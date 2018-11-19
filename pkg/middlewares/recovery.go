package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/spy16/droplets/pkg/logger"
)

// WithRecovery recovers from any panics and logs them appropriately.
func WithRecovery(logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		ri := recoveryInfo{}
		safeHandler(next, &ri).ServeHTTP(wr, req)

		if ri.panicked {
			logger.Errorf("recovered from panic: %+v", ri.val)

			wr.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(wr).Encode(map[string]interface{}{
				"error": "Something went wrong",
			})
		}
	})
}

func safeHandler(next http.Handler, ri *recoveryInfo) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		defer func() {
			if val := recover(); val != nil {
				ri.panicked = true
				ri.val = val
			}
		}()

		next.ServeHTTP(wr, req)
	})
}

type recoveryInfo struct {
	panicked bool
	val      interface{}
}
