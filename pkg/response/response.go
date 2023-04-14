package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
func JSON[T any](writer http.ResponseWriter, status int, v T) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(status)
	_, _ = writer.Write(buf.Bytes())
}
