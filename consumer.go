package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := nil amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := comm.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Messages: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")
	<-forever
}