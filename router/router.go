package router

import (
	"time"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/httpserver/app/plastik/middlewares"
	"github.com/azharprabudi/api-plastik/httpserver/app/plastik/routes"
	"github.com/go-chi/chi"
)

// InitRoute ...
func InitRoute(db *db.DB) *chi.Mux {
	// create router from chi
	r := chi.NewRouter()

	// apply middleware to all
	r.Use(middlewares.AcceptContentType)
	r.Use(middlewares.CheckContentType)
	r.Use(middlewares.CheckClientSecret)
	r.Use(middlewares.Throttle(1000))
	r.Use(middlewares.Timeout(60 * time.Second))

	// assign routes
	r.Route("/api", func(r chi.Router) {
		routes.NewRoutesV1Plastik(&r, db)
	})

	return r
}
