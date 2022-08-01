package migrations

import (
	"fmt"
	"testing"
)

type testTable struct {
	ID        string `pq:"pk"`
	FirstName string `pq:"unique,!null"`
}

func (tt testTable) Name() string {
	return "testTable"
}

func TestTable_MigrateQuery(t1 *testing.T) {
	tbl := Table[testTable]{}
	fmt.Println(tbl.MigrateQuery())
}
