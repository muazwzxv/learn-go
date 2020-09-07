package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID int `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Nickname string `json:"nickname"`
	Phone string  `json:phone_no"`
	Email string  `json:"email"`
}

func main() {
	fmt.Println("GO mysql tutorial")

	db, err := sql.Open("mysql", ":@tcp(127.0.0.1:3306)/test_crm")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "select * from users"
	result, err := db.Query(query)

	for result.Next() {
		var user User
		if err != nil {
			panic(err.Error())
		}

		err = result.Scan(&user.ID, &user.UUID, &user.Name, &user.Nickname, &user.Phone, &user.Phone, &user.Email)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)


	}
	fmt.Println(result, "Instance of sql")
}
