package main

import (
	"errors"
	"fmt"
)

type person struct {
	name string
	age  int
}
/*
func main() {
	x := make(map[string]int)

	x["muaz"] = 21
	x["afrina"] = 20
	x["shuk"] = 21

	hooman := []person{}
	var temp person
	for key, val := range x {
		temp.name = key
		temp.age = val
		hooman = append(hooman, temp)
	}

	for i := 0; i < len(hooman); i++ {
		fmt.Println(hooman[i])
	}

	keyword, err := getName("muaz")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keyword)
	}
}

func getName(name string) (string, error) {
	if name == "nuq" {
		return " ", errors.New("This is a dick")
	}

	return "Your name is " + name, nil
}
