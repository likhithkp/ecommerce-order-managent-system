package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/likhithkp/ecommerce-order-managent-system/payments/shared"
)

func MockPaymentGateway() bool {
	rand.Seed(time.Now().UnixNano())
	successRate := 95
	randomNumber := rand.Intn(100)

	return randomNumber < successRate
}

func ValidatePayment(orderData []byte) *shared.Response {
	var data shared.Order
	err := json.Unmarshal(orderData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return &shared.Response{
			Message:    "Internal server error",
			StatusCode: 500,
			Data:       false,
		}
	}

	if MockPaymentGateway() {
		fmt.Println("Payment successful for order")
		return &shared.Response{
			Message:    "Payment successful for order",
			StatusCode: 200,
			Data:       true,
		}
	} else {
		fmt.Println("Payment failed for order")
		return &shared.Response{
			Message:    "Payment failed for order",
			StatusCode: 500,
			Data:       false,
		}
	}
}
