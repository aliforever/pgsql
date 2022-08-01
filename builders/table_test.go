package builders

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"postgresql/columns"
	"postgresql/options"
	"testing"
)

type users struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

func (users) TableName() string {
	return "users"
}

func TestNewCreateTable(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable database=testapp")
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------------------
	tbl := Table[users](db)

	err = tbl.DropTable()

	err = tbl.BuildSchema().
		AddColumn(columns.NewInteger("id").PrimaryKey().Identity()).
		AddColumn(columns.NewVarchar("first_name", 20).NotNull()).
		Build()

	err = tbl.Insert(users{
		FirstName: "Ali",
	}, options.NewInsert().IgnoreFields("id"))
	if err != nil {
		panic(err)
	}

	err = tbl.Insert(users{
		FirstName: "Ali 2",
	}, options.NewInsert().IgnoreFields("id"))
	if err != nil {
		panic(err)
	}

	data, err := tbl.Query().Where("first_name", "=", "Ali").FindOne()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	// ----------------------------------------------------------------------------------
}
