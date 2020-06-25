package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO mysql tutorial")

	db, err := sql.Open("mysql", ":@tcp(127.0.0.1:3306)/test_crm")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "select * from users"
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result, "Instance of sql")
}
