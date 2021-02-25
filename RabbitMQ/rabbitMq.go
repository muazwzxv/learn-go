package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ + Go")
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	// opens a channel to RabbitMQ instance
	ch, err := con.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()

	// interact with the instance and declare a queue
	// we can publish and subscribe to
	queue, err := ch.QueueDeclare(
		"test", // name
		true,   // durable
		false,  // exclusive
		false,  // autoDelete
		false,  //
		nil,
	)
	fmt.Println(queue)

	fmt.Println("Successfully connected to RabbitMQ instance")
}
