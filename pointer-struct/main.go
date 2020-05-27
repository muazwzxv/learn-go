package main

import(
	"fmt"
	"encoding/json"
)

/*
Struct type using default data types
*/
type Book struct {
	Name string
	Pages int
	Author Author
}

type Author struct {
	Name string
	Age int
	Nationality string
	Salary float64
}

/*
struct using json encode
*/
type BookJson struct {
	Name string `json: "name"`
	Pages int `json: "pages"`
	Author AuthorJson `json: "author"`
}

type AuthorJson struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Nationality string `json:"nationality"`
	Salary float64 `json:"nationality"`
}

// Manipulates the json struct
func useJson() {
	e := Book {
		name: "The world of dicks",
		pages: 300,
		author: Author {
			name: "Muaz Bin Wazir",
			age: 21,
			nationality: "Malaysian",
			salary: 60000.00,
		},
	}
}

// Manipulates the default data type struct
func defaultDataType() {
	e := Book {
		name: "The world of dicks",
		pages: 300,
		author: Author {
			name: "Muaz Bin Wazir",
			age: 21,
			nationality: "Malaysian",
			salary: 60000.00,
		},
	}


	fmt.Println("Proper output")
	fmt.Println("Name :", e.name)
	fmt.Println("Pages :", e.pages)
	fmt.Println("Authors Name :", e.author.name)
	fmt.Println("Authos age :",e.author.age)
	fmt.Println("Authors nationality :", e.author.nationality)
	fmt.Println("Authors salary :", e.author.salary)

	fmt.Println("Display as struct")
	fmt.Println(e)
}


