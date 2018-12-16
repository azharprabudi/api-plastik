package cmd

import (
	"fmt"
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
		port = config.Port
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		panic(err)
	}
}
