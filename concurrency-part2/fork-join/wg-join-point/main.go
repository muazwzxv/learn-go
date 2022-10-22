package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()
	fmt.Println("done waiting, main exits")
	fmt.Println("elapsed: ", time.Since(now))
}

func work() {
	fmt.Println("printing something ")
}
