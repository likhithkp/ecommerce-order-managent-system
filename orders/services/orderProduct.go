package services

import (
	"encoding/json"
	"log"

	"github.com/likhithkp/ecommerce-order-managent-system/orders/shared"
)

func OrderProduct(order *shared.Order) (bool, *shared.Response) {
	isValidated, res := ValidateProduct(order)

	if isValidated {
		orderByteData, err := json.Marshal(order)

		if err != nil {
			log.Println("Error while marshaling the order")
			return true, &shared.Response{
				Message:    "Something went wrong!",
				StatusCode: 500,
				Data:       nil,
			}
		}

		res := PushOrderToKafka("order.created", "order_data", orderByteData, "localhost:9092")

		if res != nil {
			return true, res
		}

	}

	return isValidated, res
}
