package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/spy16/droplets/internal/delivery/rest"
	"github.com/spy16/droplets/internal/delivery/web"
	"github.com/spy16/droplets/internal/stores"
	"github.com/spy16/droplets/internal/usecases/users"
	"github.com/spy16/droplets/pkg/graceful"
	"github.com/spy16/droplets/pkg/logger"
	"github.com/spy16/droplets/pkg/middlewares"
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

	webHandler := web.New(web.Config{})

	router := mux.NewRouter()
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", restHandler))
	router.PathPrefix("/").Handler(webHandler)

	srv := server(lg, router)
	lg.Infof("listening for requests on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		lg.Errorf("http server exited: %s", err)
	}
}

func server(lg logger.Logger, handler http.Handler) *graceful.Server {
	handler = withMiddlewares(handler, lg)

	srv := graceful.NewServer(handler, os.Interrupt)
	srv.Addr = ":8080"
	srv.Log = lg.Errorf

	return srv
}

func withMiddlewares(handler http.Handler, logger logger.Logger) http.Handler {
	handler = middlewares.WithRequestLogging(logger, handler)
	handler = middlewares.WithRecovery(logger, handler)
	return handler
}
