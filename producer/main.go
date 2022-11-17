package main

import (
	"context"
	"fmt"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafkago.Writer{
		Addr:     kafkago.TCP("localhost:8097", "localhost:8098", "localhost:8099"),
		Topic:    "p2",
		Balancer: &kafkago.RoundRobin{},
	}

	for i := 0; i < 100; i++ {
		err := writer.WriteMessages(context.Background(), kafkago.Message{
			Value: []byte(fmt.Sprintf("hola %d", i)),
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Message successfully send!")
	}
}
