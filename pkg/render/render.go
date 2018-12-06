package render

import (
	"encoding/json"
	"io"
	"net/http"
)

const contentTypeJSON = "application/json; charset=utf-8"

// JSON uses the json encoder and writes JSON encoded version of the 'val'
// to the writer. If the writer implements http.ResponseWriter, this will
// also set appropriate content type header.
func JSON(wr io.Writer, val interface{}) error {
	if hw, ok := wr.(http.ResponseWriter); ok {
		hw.Header().Set("Content-type", contentTypeJSON)
	}

	return json.NewEncoder(wr).Encode(val)
}
