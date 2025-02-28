package services

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaclient "github.com/likhithkp/ecommerce-order-managent-system/payments/clients/kafka"
	"github.com/likhithkp/ecommerce-order-managent-system/payments/shared"
)

func PushPaymentValidationToKafka(topic, key string, byteData []byte, host string) *shared.Response {
	p := kafkaclient.CreateProducer(host)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(key),
		Value: byteData,
	}, nil)

	if err != nil {
		log.Println("Error while creating partition", err.Error())
		return &shared.Response{
			Message:    "Something went wrong!",
			StatusCode: 500,
			Data:       nil,
		}
	}

	p.Flush(500)
	return nil
}
