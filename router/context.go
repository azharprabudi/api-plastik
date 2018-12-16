package router

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/go-chi/chi"
)

// Router ...
type Router struct {
	db *db.DB
	r  *chi.Mux
}
