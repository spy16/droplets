package rest

import (
	"encoding/json"
	"net/http"

	"github.com/spy16/droplets/pkg/errors"
)

func writeResponse(wr http.ResponseWriter, status int, v interface{}) {
	wr.Header().Set("Content-Type", "application/json; charset: utf-8")
	wr.WriteHeader(status)
	if err := json.NewEncoder(wr).Encode(v); err != nil {
		if loggable, ok := wr.(errorLogger); ok {
			loggable.Errorf("failed to write data to http ResponseWriter: %s", err)
		}
	}
}

func writeError(wr http.ResponseWriter, err error) {
	if e, ok := err.(*errors.Error); ok {
		writeResponse(wr, e.Code, e)
		return
	}
	writeResponse(wr, http.StatusInternalServerError, err)
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
