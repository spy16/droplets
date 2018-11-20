package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spy16/droplets/pkg/logger"
)

// New initializes a new webapp server.
func New(lg logger.Logger, cfg Config) http.Handler {
	fsServer := newSafeFileSystemServer(lg, cfg.StaticDir)

	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", fsServer))
	router.Handle("/favicon.ico", fsServer)

	return router
}

// Config represents server configuration.
type Config struct {
	TemplateDir string
	StaticDir   string
}
