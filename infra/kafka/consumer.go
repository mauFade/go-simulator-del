package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChannel chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: msgChan,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	consumer, err := ckafka.NewConsumer(configMap)

	if err != nil {
		log.Fatalf(err.Error())
	}

	topics := []string{os.Getenv("KafkareadTopic")}

	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer started.")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			// message channel do Go recebe a mensagem do kafka
			k.MessageChannel <- msg
		}
	}
}
