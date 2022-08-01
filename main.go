package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	res, err := db.Exec("CREATE DATABASE testapp")
	if err != nil {
		fmt.Println("here")
		panic(err)
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(db.Stats())

}
