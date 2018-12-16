package qb

import "github.com/azharprabudi/api-plastik/helper/querybuilder/model"

// QueryBuilderInterface ...
type QueryBuilderInterface interface {
	Query(string, int, int, []*qbmodel.Order) string
	QueryWhere(string, []*qbmodel.Condition, []*qbmodel.Order) string
	QueryWith(string, []*qbmodel.Join, []*qbmodel.Order) string
	QueryWhereWith(string, []*qbmodel.Join, []*qbmodel.Condition, []*qbmodel.Order) string
	Create(string, interface{}) string
	Update(string, interface{}, []*qbmodel.Condition) string
	UpdateWhere(string, interface{}, []*qbmodel.Condition) string
	Delete(string, []*qbmodel.Condition) string
}
