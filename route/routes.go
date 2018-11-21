package route

import (
	"github.com/api-plastik/db"
	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
)

// InitRoute ...
func InitRoute(db *db.DB) *chi.Mux {
	// create router from chi
	r := chi.NewRouter()

	// apply middleware to all
	r.Use(middleware.AllowContentType("application/json"))

	// assign routes
	r.Route("/api", func(r chi.Router) {
		NewRoutesV1(&r, db)
	})

	return r
}
