package main

import (
	"fmt"
	"time"
)

func timely(value int) {
	for i := 0 ; i < value ; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// sequential execution
	go timely(10)
    go timely(10)
	
	// Go anonymous function
	go func() {
		word := "Keeping the concurrency running"

		for i:= 0 ; i < 20 ; i++ {
			fmt.Println(word, i)
		}
	}()
	// We are expecting an input to terminate the program
	var input string
	fmt.Scanln(&input)
}
