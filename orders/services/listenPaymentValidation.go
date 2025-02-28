package services

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaclient "github.com/likhithkp/ecommerce-order-managent-system/orders/clients/kafka"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/shared"
)

var PaymentResponseChannel = make(chan shared.Response, 1)

func ListenPaymentValidationEvent(host string, groupId string, topic string) {
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

			var res shared.Response
			json.Unmarshal(ev.Value, &res)

			PaymentResponseChannel <- res
		}
	}
}
