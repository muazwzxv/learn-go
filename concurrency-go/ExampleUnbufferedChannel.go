package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	data := rand.Intn(100)
	fmt.Println("Calculated Random Value :", data)
	time.Sleep(1000 * time.Milisecond)
	c <- data
	fmt.Println("Only executes after another goroutine performs receive on the channel")
}

func main() {
	fmt.Println("Go channel testrun")
	
	// use make function to create a channel
	channel := make(chan int) // created a channel of type int
	defer close(channel)

	go CalculateValue(channel)
	go CalculateValue(channel)

	data := <-channel
	fmt.Println("Result returned from Channel :", data)
}


