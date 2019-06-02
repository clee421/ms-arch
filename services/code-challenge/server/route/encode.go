package route

import (
	"encoding/json"
	"net/http"
)

func encode(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json, charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func encodeError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	encode(w, struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
}
