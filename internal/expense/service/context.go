package service

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/expense/command"
	"github.com/azharprabudi/api-plastik/internal/expense/query"
	"github.com/azharprabudi/api-plastik/internal/expense/transform"
)

// ExpenseService ...
type ExpenseService struct {
	db        *db.DB
	query     query.ExpenseQueryInterface
	command   command.ExpenseCommandInterface
	transform transform.ExpenseTransformInterface
}
