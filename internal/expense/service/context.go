package service

import (
	"github.com/api-plastik/internal/expense/command"
	"github.com/api-plastik/internal/expense/query"
	"github.com/api-plastik/internal/expense/transform"
)

// ExpenseService ...
type ExpenseService struct {
	query     query.ExpenseQueryInterface
	command   command.ExpenseCommandInterface
	transform transform.ExpenseTransformInterface
}
