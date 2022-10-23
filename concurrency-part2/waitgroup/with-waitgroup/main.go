package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	now := time.Now()

	processes := 10
	wg.Add(processes)

	for i := 1; i <= processes; i++ {
		go work(&wg, now, i)
	}

	wg.Wait()

	fmt.Printf("\n main is done in %v", time.Since(now))
}

func work(wg *sync.WaitGroup, now time.Time, id int) {
	defer wg.Done()
	fmt.Printf("\n Task %d is done in %v", id, time.Since(now))
}
