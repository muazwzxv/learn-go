package main

import (
	"fmt"
	"sync"
	"time"
)

func display(value int, wg *sync.WaitGroup) {
	for i := 0 ; i < value ; i++ {
		fmt.Println("Inside a fooking gorouting", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start sequence")

	var wg sync.WaitGroup
	wg.Add(1)
	go display(10, &wg)
	wg.Wait()

	fmt.Println("Sequence finished")
}
