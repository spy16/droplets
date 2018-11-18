package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spy16/droplet/pkg/errors"
	"github.com/spy16/droplet/pkg/logger"
	"github.com/spy16/droplet/pkg/middlewares"
)

// New initializes the server with routes exposing the given usecases.
func New(logger logger.Logger, reg registration, ret retriever) *Server {
	srv := &Server{}
	srv.Logger = logger

	// setup router with default handlers
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(srv.notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(srv.methodNotAllowedHandler)

	// setup api endpoints
	router.HandleFunc("/health", srv.healthCheckHandler)
	addUsersAPI(logger, router, reg, ret)

	// setup middlewares
	srv.router = middlewares.WithRequestLogging(logger, router)
	srv.router = middlewares.WithRecovery(logger, srv.router)
	return srv
}

// Server represents a REST API server.
type Server struct {
	logger.Logger

	router http.Handler
}

func (srv *Server) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(wr, req)
}

func (srv *Server) healthCheckHandler(wr http.ResponseWriter, req *http.Request) {
	info := map[string]interface{}{
		"status": "ok",
	}
	writeResponse(wr, http.StatusOK, info)
}

func (srv *Server) notFoundHandler(wr http.ResponseWriter, req *http.Request) {
	writeResponse(wr, http.StatusNotFound, errors.ResourceNotFound("path", req.URL.Path))
}

func (srv *Server) methodNotAllowedHandler(wr http.ResponseWriter, req *http.Request) {
	writeResponse(wr, http.StatusMethodNotAllowed, errors.ResourceNotFound("path", req.URL.Path))
}
