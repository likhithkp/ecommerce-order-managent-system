package services

import (
	"context"
	"fmt"

	"github.com/likhithkp/ecommerce-order-managent-system/inventory/clients/redis"
	"github.com/likhithkp/ecommerce-order-managent-system/inventory/db/repository"
	"github.com/likhithkp/ecommerce-order-managent-system/inventory/shared"
)

func AddProduct(item *shared.Item) *shared.Response {
	query := `INSERT INTO items (name, description, price, count, category) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, price, count, category`

	data, err := repository.InsertItem(query, item)

	if err != nil {
		return err
	}

	client := redis.Redis()
	key := fmt.Sprintf("item:%v", data.Id)
	redisErr := client.HSet(context.Background(), key, map[string]any{
		"itemId":          data.Id,
		"itemName":        data.Name,
		"itemDescription": data.Description,
		"itemCount":       data.Count,
		"itemPrice":       data.Price,
		"itemCategory":    data.Category,
	}).Err()

	if redisErr != nil {
		return &shared.Response{
			Message:    "Failed to cache item in Redis",
			StatusCode: 500,
			Data:       data,
		}
	}

	return &shared.Response{
		Message:    data.Name + " added successfully",
		StatusCode: 200,
		Data:       data,
	}

}
