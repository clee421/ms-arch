package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DefaultRoute if no other route works
func DefaultRoute(r *mux.Router) {
	// Default
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	encode(w, struct {
		Error string `json:"error"`
	}{
		Error: "Path not found!",
	})
}
