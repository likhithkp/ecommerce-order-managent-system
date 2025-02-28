package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/clients/redis"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/db"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/routes"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/services"
)

func main() {
	mux := http.NewServeMux()

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Unable to load .env", err.Error())
		return
	}

	redisClient := redis.Redis()
	defer redisClient.Close()

	db.ConnectDb()
	defer db.CloseDB()

	go services.ListenPaymentValidationEvent("localhost:9092", "order_validation_group", "order.validate")

	routes.InventoryRouter(mux)
	http.ListenAndServe(":3001", mux)
}
