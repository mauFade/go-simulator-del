package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/mauFade/go-simulator-del/app/route"
	"github.com/mauFade/go-simulator-del/infra/kafka"
)

func ProduceMessage(message *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()

	json.Unmarshal(message.Value, &route)

	route.LoadPositions()

	positions, err := route.ExportJsonPositions()

	if err != nil {
		log.Println(err.Error())
	}

	for _, posit := range positions {
		kafka.Publish(posit, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500) // A cada 500ms manda uma posição
	}
}
