package main

import (
	"fmt"
	"sync"
)

type request func()

func main() {
	requests := map[int]request{}

	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request: ", n)
			}
		}

		requests[i] = f(i)
	}
	fmt.Println(requests)

	var wg sync.WaitGroup
	max := 10

	fmt.Printf("\n before execution: \n %v", requests)
	loop := len(requests)
	for i := 0; i < loop; i += max {
		for j := i; j < (i + max); j++ {
			wg.Add(1)
			go func(f request) {
				defer wg.Done()
				f()
			}(requests[j+1])
			delete(requests, j+1)
		}
		wg.Wait()
		fmt.Println(max, "request processed")
	}

	fmt.Printf("\n after execution: \n %v", requests)
}
