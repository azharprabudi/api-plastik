package router

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/httpserver/app/plastik/middlewares"
	"github.com/api-plastik/httpserver/app/plastik/routes"
	"github.com/go-chi/chi"
)

// InitRoute ...
func InitRoute(db *db.DB) *chi.Mux {
	// create router from chi
	r := chi.NewRouter()

	// apply middleware to all
	r.Use(middlewares.SetContentType)
	r.Use(middlewares.CheckClientSecret)

	// assign routes
	r.Route("/api", func(r chi.Router) {
		routes.NewRoutesV1Plastik(&r, db)
	})

	return r
}