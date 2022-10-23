package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go run(wg)
	wg.Wait()
}

func run(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goroutine running")
}
