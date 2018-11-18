package main

import (
	"os"

	"github.com/spf13/viper"
	"github.com/spy16/droplet/internal/delivery/rest"
	"github.com/spy16/droplet/internal/stores"
	"github.com/spy16/droplet/internal/usecases/users"
	"github.com/spy16/droplet/pkg/graceful"
	"github.com/spy16/droplet/pkg/logger"
	"gopkg.in/mgo.v2"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("MONGO_URI", "mongodb://localhost")
	lg := logger.New(os.Stderr, "debug", "text")

	session, err := mgo.Dial(viper.GetString("MONGO_URI"))
	if err != nil {
			panic(err)
	}
	defer session.Close()

	lg.Debugf("setting up rest api service")
	userStore := stores.NewUsers(session.DB("droplets"))
	userRegistration := users.NewRegistration(lg, userStore)
	userRetriever := users.NewRetriever(lg, userStore)
	restHandler := rest.New(lg, userRegistration, userRetriever)

	srv := graceful.NewServer(restHandler, os.Interrupt)
	srv.Addr = ":8080"
	srv.Log = lg.Errorf

	lg.Infof("REST API server listening on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		lg.Errorf("http server exited: %s", err)
	}
}
