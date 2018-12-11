package qb

import "github.com/api-plastik/helper/querybuilder/model"

// QueryBuilderInterface ...
type QueryBuilderInterface interface {
	Query(string, int, int) string
	QueryWhere(string, []*model.Condition, []*model.Order) string
	QueryWith(string, []*model.Join, []*model.Order) string
	QueryWhereWith(string, []*model.Join, []*model.Condition, []*model.Order) string
	Create(string, interface{}) string
	Update(string, interface{}, []*model.Condition) string
	UpdateWhere(string, interface{}, []*model.Condition) string
	Delete(string, []*model.Condition) string
}
