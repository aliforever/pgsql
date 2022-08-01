package queries

import "fmt"

func CreateDatabase(dbName string) string {
	return fmt.Sprintf("CREATE DATABASE %s", dbName)
}

func DropDatabase(dbName string) string {
	return fmt.Sprintf("DROP DATABASE %s", dbName)
}
