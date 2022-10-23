package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()
}
