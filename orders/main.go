package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/clients/redis"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/db"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/routes"
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

	routes.InventoryRouter(mux)
	http.ListenAndServe(":3001", mux)
}
