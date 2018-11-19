package config

import (
	"net/http"

	"github.com/go-chi/chi"
)

// StartServer ...
func StartServer(port string, r *chi.Mux) {
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
