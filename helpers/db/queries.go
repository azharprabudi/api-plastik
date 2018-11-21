package helpers

import (
	"fmt"
	"reflect"

	"github.com/api-plastik/helpers/db/model"
)

// this function below, to decide is the last loop or not
func isLastIteration(currNumb int, endNumb int) bool {
	if (currNumb + 1) == endNumb {
		return true
	}
	return false
}

// check reflect kind is string or not
func isString(dType reflect.Kind) bool {
	if dType == reflect.String {
		return true
	}
	return false
}

// check reflect kind is int or not
func isInt(dType reflect.Kind) bool {
	if dType == reflect.Int || dType == reflect.Int8 || dType == reflect.Int16 || dType == reflect.Int32 || dType == reflect.Int64 {
		return true
	}
	return false
}

// check reflect kind is float or not
func isFloat(dType reflect.Kind) bool {
	if dType == reflect.Float32 || dType == reflect.Float64 {
		return true
	}
	return false
}

// check reflect kind is bool or not
func isBool(dType reflect.Kind) bool {
	if dType == reflect.Bool {
		return true
	}
	return false
}

func createQueriesInsert(data interface{}) (string, string) {
	var cols string
	var values string

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		var separComma string
		if isLastIteration(i, v.NumField()) == true {
			separComma = ""
		} else {
			separComma = ","
		}

		// set column
		columnName := v.Type().Field(i).Tag.Get("db")
		cols = fmt.Sprintf("%s\"%s%s \"", cols, columnName, separComma)

		// set value
		value := v.Field(i)
		dTypeVal := value.Kind()

		// check value by data type
		if isString(dTypeVal) == true {
			values = fmt.Sprintf("%s'%s'%s ", values, value, separComma)
		} else {
			values = fmt.Sprintf("%s%s%s ", values, value, separComma)

		}
	}
	return cols, values

}

func createQueriesUpdate(data interface{}) string {
	var sets string

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		separComma := ","
		if isLastIteration(i, v.NumField()) == true {
			separComma = ""
		}
		// get column
		column := v.Type().Field(i).Tag.Get("db")

		// get value
		value := v.Field(i)
		dTypeVal := value.Kind()

		// check value by data type
		if isString(dTypeVal) == true {
			sets = fmt.Sprintf("%s%s='%s'%s", sets, column, value, separComma)
		} else {
			sets = fmt.Sprintf("%s%s=%v%s", sets, column, value, separComma)

		}
	}
	return sets
}

func createQueriesWhere(conditions []*model.Condition) string {
	var where string
	for index, condition := range conditions {
		nextCondition := condition.NextCond
		if isLastIteration(index, len(conditions)) {
			nextCondition = ""
		}

		if isString(reflect.TypeOf(condition.Value).Kind()) == true {
			where = fmt.Sprintf("%s \"%s\"%s'%v' %s", where, condition.Key, condition.Operator, condition.Value, nextCondition)
		} else {
			where = fmt.Sprintf("%s \"%s\"%s%v %s", where, condition.Key, condition.Operator, condition.Value, nextCondition)
		}
	}
	return where
}

func createQueriesWith(joins []*model.Join) string {
	var with string
	for _, join := range joins {
		with = fmt.Sprintf(" %s %v %s %s ON %s=%s", with, join.TableFrom, join.Type, join.TableWith, join.ColumnTableFrom, join.ColumnTableWith)
	}
	return with
}
