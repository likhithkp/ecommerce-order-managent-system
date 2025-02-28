package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/likhithkp/ecommerce-order-managent-system/orders/clients/redis"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/shared"
)

func ValidateProduct(order *shared.Order) (bool, *shared.Response) {
	client := redis.Redis()

	for _, product := range order.Products {
		key := fmt.Sprintf("item:%v", product.ProdID)
		data, err := client.HGetAll(context.Background(), key).Result()
		if err != nil {
			log.Println("Error while fetching item data from redis", err.Error())
			return false, &shared.Response{
				Message:    "Unable to check the product",
				StatusCode: 500,
				Data:       nil,
			}
		}

		if len(data) == 0 {
			log.Println("No items found in redis")
			return false, &shared.Response{
				Message:    "Unable to find the item/items",
				StatusCode: 500,
				Data:       nil,
			}
		}

		itemCount, _ := strconv.Atoi(data["itemCount"])
		pricePerUnit, _ := strconv.Atoi(data["itemPrice"])

		if itemCount < product.Quantity {
			fmt.Println("Product out of stock")
			return false, &shared.Response{
				Message:    "Product out of stock",
				StatusCode: 500,
				Data:       nil,
			}
		}

		total := float64(product.Quantity * pricePerUnit)
		if order.Payment.TotalAmount != total {
			fmt.Println("Payment amount mismatch")
			return false, &shared.Response{
				Message:    "Payment amount mismatch",
				StatusCode: 500,
				Data:       nil,
			}
		}
	}
	return true, &shared.Response{
		Message:    "Order validated successfully",
		StatusCode: 200,
		Data:       nil,
	}
}
