package service

import (
	"github.com/azharprabudi/api-plastik/internal/seller/command"
	"github.com/azharprabudi/api-plastik/internal/seller/query"
	"github.com/azharprabudi/api-plastik/internal/seller/transform"
)

// SellerService ...
type SellerService struct {
	query     query.SellerQueryInterface
	command   command.SellerCommandInterface
	transform transform.SellerTransformInterface
}
