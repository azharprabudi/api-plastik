package helpers

import (
	"context"

	"github.com/api-plastik/helpers/db/model"
	"github.com/jmoiron/sqlx"
)

// DBQueryBuilderInterface ...
type DBQueryBuilderInterface interface {
	CreateTrx() context.Context
	RunTrx(context.Context, func(*sqlx.Tx) error) error
	Query(context.Context) (*sqlx.Rows, error)
	QueryWhere(context.Context, []*model.Condition) (*sqlx.Row, error)
	QueryWith(context.Context, []*model.Join) (*sqlx.Rows, error)
	QueryWhereWith(context.Context, []*model.Join, []*model.Condition) (*sqlx.Rows, error)
	RawQuery(context.Context, string) (*sqlx.Rows, error)
	Create(context.Context, interface{}, bool) error
	// Update(context.Context, string, int, int) error
	// Delete(context.Context, string, int, int) error
}
