package main

import (
	"errors"
	"fmt"
)

func main() {
	selection()

	result, err := name("nuq")

	fmt.Println("results:", result, "Error:", err)
}

func selection() {
	x := []string{"muaz", "nuq", "afrina", "lan", "danisha", "al", "alya", "shuk", "anis", "bella", "copia", "arep"}

	for index := range x {
		if x[index] == "muaz" {
			fmt.Println("Paling kacak is ", x[index])
			break
		}

		fmt.Println("Korang semua tak kacak")
	}
}

func loop() {
	fmt.Println("This is how control structure works")

	x := []string{"muaz", "nuq", "afrina", "lan", "danisha", "al", "alya", "shuk", "anis", "bella", "copia", "arep"}

	for index, val := range x {
		fmt.Println("index:", index, "value: "+val)
	}
}

func name(name string) (string, error) {
	if name != "muaz" {
		return "you are not muaz", errors.New("fuck off")
	}

	return "kacak thialll", nil
}


