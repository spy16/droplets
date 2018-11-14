package main

import (
	"os"

	"github.com/spy16/droplet/internal/delivery/rest"
	"github.com/spy16/droplet/pkg/graceful"
	"github.com/spy16/droplet/pkg/logger"
)

func main() {
	lg := logger.New(os.Stderr, "debug", "text")

	lg.Debugf("setting up rest api service")
	srv := graceful.NewServer(rest.New(lg), os.Interrupt)
	srv.Addr = ":8080"
	srv.Log = lg.Errorf

	lg.Infof("REST API server listening on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		lg.Errorf("http server exited: %s", err)
	}
}
