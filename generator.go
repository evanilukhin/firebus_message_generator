package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/satori/go.uuid"
	"time"
)

//Message - message sended to topic
type Message struct {
	Item      int    `json:"item"`
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}

func main() {
	kafkaBrokerHost := flag.String("kafka", "0.0.0.0:9092", "Kafka broker address. Default: 0.0.0.0:9092")
	topic := flag.String("topic", "firebus", "Destination topic. Default: firebus")
	countMessages := flag.Int("count", 100, "Count messages in a group. Default: 100")

	flag.Parse()

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *kafkaBrokerHost, "go.delivery.reports": false})
	if err != nil {
		panic(err)
	}

	uuid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Start generating")
	for i := 0; i < *countMessages; i++ {
		message := Message{
			UUID:      uuid,
			Item:      i,
			CreatedAt: time.Now().UTC().Format("15:04:05.999999"),
		}
		marshalledMessage, _ := json.Marshal(message)

		kafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
			Value:          marshalledMessage,
		}, nil)
		if i%1000000 == 0 {
			fmt.Printf("Sended %d messages\n", i)
		}
	}
	kafkaProducer.Flush(30000)
}
