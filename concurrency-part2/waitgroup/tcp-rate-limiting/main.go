package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3

	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		limit := max
		if (i + max) > total {
			limit = total - i
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(n int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("failed to connect: %v", err)
				}

				stream, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("failed to read stream: %v", err)
				}

				if string(stream) != "success" {
					log.Fatalf("request error, request : %d", i+1+j)
				}

				fmt.Printf("request %d: success\n", i+1+j)
			}(j)
		}
		wg.Wait()
	}
}
