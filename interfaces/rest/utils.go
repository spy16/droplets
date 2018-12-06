package rest

import (
	"encoding/json"
	"net/http"

	"github.com/spy16/droplets/pkg/errors"
	"github.com/spy16/droplets/pkg/render"
)

func respond(wr http.ResponseWriter, status int, v interface{}) {
	if err := render.JSON(wr, status, v); err != nil {
		if loggable, ok := wr.(errorLogger); ok {
			loggable.Errorf("failed to write data to http ResponseWriter: %s", err)
		}
	}
}

func respondErr(wr http.ResponseWriter, err error) {
	if e, ok := err.(*errors.Error); ok {
		respond(wr, e.Code, e)
		return
	}
	respond(wr, http.StatusInternalServerError, err)
}

func readRequest(req *http.Request, v interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(v); err != nil {
		return errors.Validation("Failed to read request body")
	}

	return nil
}

type errorLogger interface {
	Errorf(msg string, args ...interface{})
}
