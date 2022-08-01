package migrations

import (
	"reflect"
)

type table interface {
	Name() string
}

type Table[T table] struct {
}

func (t Table[T]) processFields() {
	var table T

	tType := reflect.TypeOf(table)

	for i := 0; i < tType.NumField(); i++ {
		tag := tType.Field(i).Tag
	}
}

func (t Table[T]) CreateQuery() string {

	return `CREATE TABLE [IF NOT EXISTS] %s ()`
}
