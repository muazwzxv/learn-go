package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role Role   `json:"role"`
}

type Role struct {
	Title string `json:"title"`
	Level int    `json:"level"`
}

func main() {
	role := Role{Title: "Software Engineer", Level: 5}
	Person := Person{Name: "Muaz Bin Wazir", Age: 21, Role: role}

	byteArray, err := json.Marshal(Person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
