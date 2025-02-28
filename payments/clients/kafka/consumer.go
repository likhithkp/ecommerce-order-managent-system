package kafkaclient

import (
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	Consumer *kafka.Consumer
	once     sync.Once
)

func CreateConsumer(host string, groupId string) *kafka.Consumer {
	once.Do(func() {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"group.id":          groupId,
			"auto.offset.reset": "earliest"})

		if err != nil {
			log.Fatalln("Error while creating consumer", err.Error())
		}
		Consumer = c
	})
	return Consumer
}
