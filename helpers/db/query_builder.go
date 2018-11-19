package helpers

import (
	"fmt"
)

// Query ...
func Query(tableName string, limit int, offset int) string {
	var query string
	if limit == 0 && offset == 0 {
		query = fmt.Sprintf("select * from %s", tableName)
	} else if limit != 0 && offset == 0 {
		query = fmt.Sprintf("select * from %s limit %d", tableName, limit)
	} else {
		query = fmt.Sprintf("select * from %s limit %d offset %d", tableName, limit, offset)
	}
	return query
}

// QueryFilter ...
// func QueryFilter(tableName string, conditional interface{}) string {

// }
