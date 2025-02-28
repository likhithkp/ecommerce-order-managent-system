package services

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaclient "github.com/likhithkp/ecommerce-order-managent-system/payments/clients/kafka"
)

func OrderConsumer(host string, groupId string, topic string) {
	c := kafkaclient.CreateConsumer(host, groupId)
	defer c.Close()

	if err := c.Subscribe(topic, nil); err != nil {
		log.Println("Error while subscribing to order.created", err.Error())
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(1 * time.Second)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("Consumer error: %v\n", err)
				continue
			}

			log.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			res := ValidatePayment(ev.Value)

			data, err := json.Marshal(&res)
			if err != nil {
				log.Fatalln("Error while marshaling validation response")
			}
			PushPaymentValidationToKafka("order.validate", "order_validation", data, "localhost:9092")
		}
	}
}
