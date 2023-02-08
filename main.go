package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/mauFade/go-simulator-del/app/kafka"
	"github.com/mauFade/go-simulator-del/infra/kafka"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)

	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.ProduceMessage(msg)
		fmt.Println(string(msg.Value))
	}

}
