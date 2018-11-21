package route

import (
	"github.com/api-plastik/db"
	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
)

// InitRoute ...
func InitRoute(db *db.DB) *chi.Mux {
	// initialize route abstract
	route := new(Route)

	// create router from chi
	r := chi.NewRouter()

	// assign value to route struct
	route.r = r
	route.db = db

	// apply middleware to all
	r.Use(middleware.AllowContentType("application/json"))

	// assign routes
	NewRoutesV1(route)

	return route.r
}
