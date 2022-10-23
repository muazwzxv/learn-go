package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		run()
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("outside goroutine")
	}

	defer func() {
		fmt.Println("execution finished")
	}()

}

func run() {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("inside goroutine")
	}
}
