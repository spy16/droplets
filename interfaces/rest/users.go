package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/logger"
)

func addUsersAPI(router *mux.Router, reg registration, ret retriever, logger logger.Logger) {
	uc := &userController{
		Logger: logger,
		reg:    reg,
		ret:    ret,
	}

	router.HandleFunc("/v1/users/{name}", uc.get).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/", uc.search).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/", uc.post).Methods(http.MethodPost)
}

type userController struct {
	logger.Logger
	reg registration
	ret retriever
}

func (uc *userController) get(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user, err := uc.ret.Get(req.Context(), vars["name"])
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, user)
}

func (uc *userController) search(wr http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()["t"]
	users, err := uc.ret.Search(req.Context(), vals, 10)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, users)
}

func (uc *userController) post(wr http.ResponseWriter, req *http.Request) {
	user := domain.User{}
	if err := readRequest(req, &user); err != nil {
		uc.Warnf("failed to read user request: %s", err)
		respond(wr, http.StatusBadRequest, err)
		return
	}

	registered, err := uc.reg.Register(req.Context(), user)
	if err != nil {
		uc.Warnf("failed to register user: %s", err)
		respondErr(wr, err)
		return
	}

	uc.Infof("new user registered with id '%s'", registered.Name)
	respond(wr, http.StatusCreated, registered)
}

type registration interface {
	Register(ctx context.Context, user domain.User) (*domain.User, error)
}

type retriever interface {
	Get(ctx context.Context, name string) (*domain.User, error)
	Search(ctx context.Context, tags []string, limit int) ([]domain.User, error)
	VerifySecret(ctx context.Context, name, secret string) bool
}
