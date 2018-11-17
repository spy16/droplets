package main

import (
	"context"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spy16/droplet/internal/delivery/rest"
	"github.com/spy16/droplet/internal/stores"
	"github.com/spy16/droplet/internal/usecases/users"
	"github.com/spy16/droplet/pkg/graceful"
	"github.com/spy16/droplet/pkg/logger"
)

func main() {
	lg := logger.New(os.Stderr, "debug", "text")

	client, err := mongo.NewClient("mongodb://localhost")
	if err != nil {
		lg.Errorf("failed to setup MongoDB client: %s", err)
		os.Exit(1)
	}

	if err := client.Connect(context.Background()); err != nil {
		lg.Errorf("failed to setup MongoDB client: %s", err)
		os.Exit(1)
	}

	lg.Debugf("setting up rest api service")
	userStore := stores.NewUsers(client.Database("droplet"))
	userRegistration := users.NewRegistration(userStore)
	restHandler := rest.New(lg, userRegistration)

	srv := graceful.NewServer(restHandler, os.Interrupt)
	srv.Addr = ":8080"
	srv.Log = lg.Errorf

	lg.Infof("REST API server listening on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		lg.Errorf("http server exited: %s", err)
	}
}
