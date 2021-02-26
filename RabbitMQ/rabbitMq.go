package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Tutorial from https://tutorialedge.net/golang/go-rabbitmq-tutorial/
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
		false,
		nil,
	)
	fmt.Println(queue) // Print the queue status
	if err != nil {
		fmt.Println(err)
	}

	// Publish a message to the queue!
	err = ch.Publish(
		"",
		"test",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Muaz paling terhensem"),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully connected to RabbitMQ instance")
}
