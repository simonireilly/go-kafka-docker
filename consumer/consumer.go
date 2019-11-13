package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Booting Consumer")

	topicPrefix := os.Getenv("KAFKA_TOPIC_PREFIX")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"group.id":             "myGroup",
		"auto.offset.reset":    "earliest",
		"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
		"sasl.username"    : os.Getenv("KAFKA_USERNAME"),
		"sasl.password"    : os.Getenv("KAFKA_PASSWORD"),
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms"  : "SCRAM-SHA-256",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topicPrefix + "myTopic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
