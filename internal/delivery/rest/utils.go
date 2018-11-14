package rest

import (
	"encoding/json"
	"net/http"

	"github.com/spy16/droplet/pkg/logger"
)

func writeJSON(wr http.ResponseWriter, status int, v interface{}, logger logger.Logger) {
	wr.Header().Set("Content-Type", "application/json; charset: utf-8")
	wr.WriteHeader(status)
	if err := json.NewEncoder(wr).Encode(v); err != nil {
		logger.Errorf("failed to write data to http ResponseWriter: %s", err)
	}
}
