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
	json_string := `
	{
		"name" : "world of dicks 2",
		"pages": "302",
		"author ": 
	}`
}

// Manipulates the default data type struct
func defaultDataType() {
	author := Author {
		Name: "Muaz Bin Wazir",
		Age: 21,
		Nationality: "Malaysian",
		Salary: 60000.00,
	}

	e := Book {
		Name: "The world of dicks",
		Pages: 300,
		Author: author,
	}

	fmt.Println("Proper output")
	fmt.Println("Name :", e.Name)
	fmt.Println("Pages :", e.Pages)
	fmt.Println("Authors Name :", e.Author.Name)
	fmt.Println("Authos age :",e.Author.Age)
	fmt.Println("Authors nationality :", e.Author.Nationality)
	fmt.Println("Authors salary :", e.Author.Salary)

	fmt.Println("Display as struct")
	fmt.Println(e)
}

// Entry point
func main() {
	defaultDataType()
}

