package cmd

import (
	"net/http"
	"os"

	"github.com/azharprabudi/api-plastik/config"
	"github.com/go-chi/chi"
)

// StartServer ...
func StartServer(r *chi.Mux) {
	// get conf port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = os.Getenv(config.Port)
	}

	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
