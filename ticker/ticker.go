package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("emitted")
	}
}

/*
* Ticker is used to perform a task repeatedly at a given interval of time
 */
func main() {
	fmt.Println("This is ticker example")

	go backgroundTask()

	fmt.Println("The rest of my application can continue")

	// To keep our main function alive indefinitely
	select {}
}
