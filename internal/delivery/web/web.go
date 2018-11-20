package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New initializes a new webapp server.
func New(cfg Config) http.Handler {
	router := mux.NewRouter()
	router.Handle(cfg.StaticDir, newSafeFileSystemServer(cfg.StaticDir))

	return router
}

// Config represents server configuration.
type Config struct {
	Addr        string
	TemplateDir string
	StaticDir   string
}
