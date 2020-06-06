package main

import (
	"fmt"
	"math/rang"
)

func CalculateValue(values chan int) {
	data:= rand.Intn(10)
	fmt.Println("Calculated Random Value {} ", values)
	values <- data
}

func main() {
	fmt.Println("Go channel testrun")

	values := make(chan int)
	defer close(values)

	go CalculateValue(values)
	value := <-values
	fmt.Println(value)
}
