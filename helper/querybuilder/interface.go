package qb

import "github.com/api-plastik/helper/querybuilder/model"

// QueryBuilderInterface ...
type QueryBuilderInterface interface {
	Query(string, int, int) string
	QueryWhere(string, []*model.Condition) string
	QueryWith(string, []*model.Join) string
	QueryWhereWith(string, []*model.Join, []*model.Condition) string
	Create(string, interface{}) string
	Update(string, interface{}, []*model.Condition) string
	Delete(string, []*model.Condition) string
}
