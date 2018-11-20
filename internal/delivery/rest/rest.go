package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/logger"
)

// New initializes the server with routes exposing the given usecases.
func New(logger logger.Logger, reg registration, ret retriever) http.Handler {
	// setup router with default handlers
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)

	// setup api endpoints
	addUsersAPI(logger, router, reg, ret)

	return router
}

func notFoundHandler(wr http.ResponseWriter, req *http.Request) {
	writeResponse(wr, http.StatusNotFound, errors.ResourceNotFound("path", req.URL.Path))
}

func methodNotAllowedHandler(wr http.ResponseWriter, req *http.Request) {
	writeResponse(wr, http.StatusMethodNotAllowed, errors.ResourceNotFound("path", req.URL.Path))
}
