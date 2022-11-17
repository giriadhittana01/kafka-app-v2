package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/namsral/flag"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

var (
	// kafka
	kafkaBrokerUrl     string
	kafkaTopic         string
	kafkaConsumerGroup string
)

func main() {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:8097,localhost:8098,localhost:8099", "Kafka brokers in comma separated value")
	flag.StringVar(&kafkaTopic, "kafka-topic", "p2", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group-a", "Kafka consumer group")

	flag.Parse()

	brokers := strings.Split(kafkaBrokerUrl, ",")

	// make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: kafkaConsumerGroup,
		Topic:   kafkaTopic,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Msgf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}
