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
	timely(10)
    timely(20)
	
	// We are expecting an input to terminate the program
	var input string
	fmt.Scanln(&input)
}
