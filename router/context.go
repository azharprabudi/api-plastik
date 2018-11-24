package router

import (
	"github.com/api-plastik/db"
	"github.com/go-chi/chi"
)

// Router ...
type Router struct {
	db *db.DB
	r  *chi.Mux
}
