package route

import (
	"github.com/api-plastik/db"
	"github.com/go-chi/chi"
)

// Route ...
type Route struct {
	db *db.DB
	r  *chi.Mux
}
