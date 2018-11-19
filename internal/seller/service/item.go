package service

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/seller/command"
	"github.com/api-plastik/internal/seller/query"
)

// NewSellerService ...
func NewSellerService(db *db.DB) SellerServiceInterface {
	return &SellerService{
		query:   query.NewSellerQuery(db),
		command: command.NewSellerCommand(db),
	}
}
