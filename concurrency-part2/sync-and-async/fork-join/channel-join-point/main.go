package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work()
		done <- struct{}{}
	}()

	<-done
	fmt.Println("done waiting, main exits")
	fmt.Println("elapsed: ", time.Since(now))
}

func work() {
	fmt.Println("printing something ")
}
