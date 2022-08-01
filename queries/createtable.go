package queries

func NewCreateTable() string {
	return `CREATE TABLE %s ()`
}

/*
CREATE TABLE %s (
	ID INT PRIMARY KEY NOT NULL
)

*/
