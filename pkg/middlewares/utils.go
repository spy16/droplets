package middlewares

import (
	"net/http"

	"github.com/spy16/droplets/pkg/logger"
)

func requestInfo(req *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"path":   req.URL.Path,
		"query":  req.URL.RawQuery,
		"method": req.Method,
		"client": req.RemoteAddr,
	}
}

func wrap(wr http.ResponseWriter, logger logger.Logger) *wrappedWriter {
	return &wrappedWriter{
		ResponseWriter: wr,
		Logger:         logger,
		wroteStatus:    http.StatusOK,
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	logger.Logger

	wroteStatus int
}

func (wr *wrappedWriter) WriteHeader(statusCode int) {
	wr.wroteStatus = statusCode
	wr.ResponseWriter.WriteHeader(statusCode)
}
