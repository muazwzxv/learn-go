package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Microsecond)
		fmt.Println("goroutine done")
	}()
	wg.Wait()
	fmt.Println("finished")
}

func deadlock() {
	// deadlock happens because we add 2 goroutines to wait but only
	// called Done() once which resulted in deadlock as waitgroup is waiting
	// for Done() that will never be called
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Microsecond)
		fmt.Println("goroutine done")
	}()
	wg.Wait()
	fmt.Println("finished")
}
