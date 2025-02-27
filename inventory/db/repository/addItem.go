package repository

import (
	"context"
	"log"

	"github.com/likhithkp/ecommerce-order-managent-system/inventory/db"
	"github.com/likhithkp/ecommerce-order-managent-system/inventory/shared"
)

func InsertItem(query string, item *shared.Item) (*shared.DbItem, *shared.Response) {
	data := new(shared.DbItem)

	err := db.DB.QueryRow(context.Background(), query, item.Name, item.Description, item.Price, item.Count, item.Category).
		Scan(&data.Id, &data.Name, &data.Description, &data.Price, &data.Count, &data.Category)

	if err != nil {
		log.Println("Error while inserting item data", err.Error())

		return nil, &shared.Response{
			Message:    "Failed to add item",
			StatusCode: 500,
			Data:       nil,
		}
	}

	return data, nil
}
