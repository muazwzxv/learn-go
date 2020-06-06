package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	data:= rand.Intn(100)
	fmt.Println("Calculated Random Value", data)
	values <- data
}

func main() {
	fmt.Println("Go channel testrun")
	
	// use make function to create a channel
	values := make(chan int)
	defer close(values)

	go CalculateValue(values)

	value := <-values
	fmt.Println("Result returned from Channel", value)
}
