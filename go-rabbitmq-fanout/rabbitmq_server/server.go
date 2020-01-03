package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://yuxxto56:123456@192.168.1.135:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//声明通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	body := fmt.Sprintf("hello:%d",time.Now().Unix())

	t := time.Now()

	err = ch.Publish(
		"order_exchange", // exchange fanout
		"",                    // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	elapsed := time.Since(t)
	failOnError(err, "Failed to publish a message")
	log.Println("[x] Send Second:",elapsed)
	log.Printf("[x] Sent %s", body)
}
