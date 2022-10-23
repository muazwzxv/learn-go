package main

import "sync"

func main() {
	var wg sync.WaitGroup

	count := 100
	for i := 1; i <= count; i++ {

		// go routine spins off with every iterations
		go func() {
			wg.Add(3)
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			wg.Wait()

			/**
			by the time all these goroutines are finished, the loop has gone to a new iterations
			which uses the wg before it is done thus causing the error that we reuse the wg before
			it is finished
			*/

		}()
	}
}
