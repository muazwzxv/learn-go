package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ + Go")
	con, err := amqp.Dial("ampq://guest:guest@localhost5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer con.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")
}
