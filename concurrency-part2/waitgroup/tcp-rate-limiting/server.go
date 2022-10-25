package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("could not create listener %v", err)
	}

	var count int32
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		count++

		go func() {
			// server connections
			defer func() {
				conn.Close()
				atomic.AddInt32(&count, -1)
			}()

			if atomic.LoadInt32(&count) > 3 {
				return
			}

			time.Sleep(100 * time.Millisecond)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Fatalf("failed to write connection %v", err)
			}
		}()
	}
}
