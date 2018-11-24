package qb

import (
	"fmt"

	"github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/helper/querybuilder/queries"
)

/*
*
* this query builder, just working on sql db. Not compatible for
* nosql db
 */

// Query ...
func (qb *QueryBuilder) Query(tableName string, limit int, offset int) string {
	query := fmt.Sprintf("select * from \"%s\"", tableName)
	if limit != 0 && offset == 0 {
		query = fmt.Sprintf("%s limit=%d", query, limit)
	} else if limit == 0 && offset != 0 {
		query = fmt.Sprintf("%s offset=%d", query, offset)
	} else if limit != 0 && offset != 0 {
		query = fmt.Sprintf("%s limit=%d offset=%d", query, limit, offset)
	}
	return query
}

// QueryWhere ...
func (qb *QueryBuilder) QueryWhere(tableName string, conditions []*model.Condition) string {
	// build query
	query := fmt.Sprintf("select * from \"%s\" where", tableName)
	where := queries.CreateQueriesWhere(conditions)
	query = fmt.Sprintf("%s%s", query, where)
	return query
}

// QueryWith ...
func (qb *QueryBuilder) QueryWith(tableName string, joins []*model.Join) string {
	// build query
	withs := queries.CreateQueriesWith(joins)
	query := fmt.Sprintf("select * from \"%s\" %s", tableName, withs)
	return query
}

// QueryWhereWith ...
func (qb *QueryBuilder) QueryWhereWith(tableName string, joins []*model.Join, conditions []*model.Condition) string {

	// build query
	withs := queries.CreateQueriesWith(joins)
	where := queries.CreateQueriesWhere(conditions)
	query := fmt.Sprintf("select * from \"%s\" %s %s", tableName, withs, where)
	return query
}

// Create ...
func (qb *QueryBuilder) Create(tableName string, data interface{}) string {
	cols, values := queries.CreateQueriesInsert(data)
	query := fmt.Sprintf("INSERT INTO \"%s\" (%s) VALUES (%s)", tableName, cols, values)
	return query
}

// Update ...
func (qb *QueryBuilder) Update(tableName string, data interface{}, conditions []*model.Condition) string {
	upd := queries.CreateQueriesUpdate(data)
	query := fmt.Sprintf("UPDATE \"%s\" %s", tableName, upd)
	return query
}

// UpdateWhere ...
func (qb *QueryBuilder) UpdateWhere(tableName string, data interface{}, conditions []*model.Condition) string {
	upd := queries.CreateQueriesUpdate(data)
	withs := queries.CreateQueriesWhere(conditions)
	query := fmt.Sprintf("UPDATE \"%s\" SET %s WHERE %s", tableName, upd, withs)
	return query
}

// Delete ...
func (qb *QueryBuilder) Delete(tableName string, conditions []*model.Condition) string {
	withs := queries.CreateQueriesWhere(conditions)
	query := fmt.Sprintf("DELETE FROM \"%s\" %s", tableName, withs)
	return query
}

// NewQueryBuilder ...
func NewQueryBuilder() QueryBuilderInterface {
	return QueryBuilderSingleton
}
