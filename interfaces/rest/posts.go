package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spy16/droplets/domain"
	"github.com/spy16/droplets/pkg/logger"
	"github.com/spy16/droplets/pkg/middlewares"
	"github.com/spy16/droplets/usecases/posts"
)

func addPostsAPI(router *mux.Router, pub postPublication, ret postRetriever, lg logger.Logger) {
	pc := &postController{}
	pc.ret = ret
	pc.pub = pub
	pc.Logger = lg

	router.HandleFunc("/v1/posts", pc.search).Methods(http.MethodGet)
	router.HandleFunc("/v1/posts/{name}", pc.get).Methods(http.MethodGet)
	router.HandleFunc("/v1/posts/{name}", pc.delete).Methods(http.MethodDelete)
	router.HandleFunc("/v1/posts", pc.post).Methods(http.MethodPost)
}

type postController struct {
	logger.Logger

	pub postPublication
	ret postRetriever
}

func (pc *postController) search(wr http.ResponseWriter, req *http.Request) {
	posts, err := pc.ret.Search(req.Context(), posts.Query{})
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, posts)
}

func (pc *postController) get(wr http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	post, err := pc.ret.Get(req.Context(), name)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, post)
}

func (pc *postController) post(wr http.ResponseWriter, req *http.Request) {
	post := domain.Post{}
	if err := readRequest(req, &post); err != nil {
		pc.Warnf("failed to read user request: %s", err)
		respond(wr, http.StatusBadRequest, err)
		return
	}
	user, _ := middlewares.User(req)
	post.Owner = user

	published, err := pc.pub.Publish(req.Context(), post)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusCreated, published)
}

func (pc *postController) delete(wr http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	post, err := pc.pub.Delete(req.Context(), name)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, post)
}

type postRetriever interface {
	Get(ctx context.Context, name string) (*domain.Post, error)
	Search(ctx context.Context, query posts.Query) ([]domain.Post, error)
}

type postPublication interface {
	Publish(ctx context.Context, post domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, name string) (*domain.Post, error)
}
