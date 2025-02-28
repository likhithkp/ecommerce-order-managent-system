package main

import (
	"net/http"

	"github.com/likhithkp/ecommerce-order-managent-system/payments/services"
)

func main() {
	go services.OrderConsumer("localhost:9092", "order_group", "order.created")
	http.ListenAndServe(":3002", nil)
}
