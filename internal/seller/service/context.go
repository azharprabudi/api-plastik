package service

import (
	"github.com/api-plastik/internal/seller/command"
	"github.com/api-plastik/internal/seller/query"
	"github.com/api-plastik/internal/seller/transform"
)

// SellerService ...
type SellerService struct {
	query     query.SellerQueryInterface
	command   command.SellerCommandInterface
	transform transform.SellerTransformInterface
}
