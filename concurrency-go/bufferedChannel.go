package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Calculate(c chan int) {
	data := rand.Intn(100)
	fmt.Println("Random value generated", data)
	time.Sleep(1000 * time.Millisecond)
	c <- data
	fmt.Println("This will execute everytime")
}

func main() {
	fmt.Println("Buffered Channel Tutorial")

	channel := make(chan int, 2)
	defer close(channel)

	go Calculate(channel)
	go Calculate(channel)

	data := <-channel
	fmt.Println("\n Return data from channel", data)

	time.Sleep(1000 * time.Millisecond)
}

func muazKacak() {
	fmt.Println("Muaz kacak rox")
}
