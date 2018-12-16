package cmd

import (
	"net/http"
	"os"

	"github.com/azharprabudi/api-plastik/config"

	"github.com/go-chi/chi"
)

// StartServer ...
func StartServer(r *chi.Mux) {
	// get the port, from env variable at server local or deployed
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = os.Getenv(config.Port)

	}

	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
