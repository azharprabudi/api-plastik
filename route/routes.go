package route

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/middlewares"
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
		NewRoutesV1(&r, db)
	})

	return r
}
