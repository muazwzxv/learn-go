package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go work(i)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n Task %d is done", id)
}
