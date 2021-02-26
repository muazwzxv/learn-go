package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go + RabbitMq consume script")
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Broker connection failed")
		panic(err)
	}

	ch, err := con.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	mssg, err := ch.Consume(
		"test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for dat := range mssg {
			fmt.Printf("Message Received: %s \n", dat.Body)
		}
	}()

	fmt.Println("Connected to RabbitMQ instance")
	fmt.Println(" [*] - Waiting for Messages")
	<-forever
}
