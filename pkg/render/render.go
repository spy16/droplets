package render

import (
	"encoding/json"
	"io"
	"net/http"
)

const contentTypeJSON = "application/json; charset=utf-8"

// JSON encodes the given val using the standard json package and writes
// the encoding output to the given writer. If the writer implements the
// http.ResponseWriter interface, then this function will also set the
// proper JSON content-type header with charset as UTF-8.
func JSON(wr io.Writer, val interface{}) error {
	if hw, ok := wr.(http.ResponseWriter); ok {
		hw.Header().Set("Content-type", contentTypeJSON)
	}

	return json.NewEncoder(wr).Encode(val)
}
