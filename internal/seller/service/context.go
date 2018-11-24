package service

import (
	"github.com/api-plastik/internal/seller/command"
	"github.com/api-plastik/internal/seller/query"
)

// SellerService ...
type SellerService struct {
	query   query.SellerQueryInterface
	command command.SellerCommandInterface
}
