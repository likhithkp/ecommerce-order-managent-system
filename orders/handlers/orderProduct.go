package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/likhithkp/ecommerce-order-managent-system/orders/services"
	"github.com/likhithkp/ecommerce-order-managent-system/orders/shared"
)

func OrderProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := shared.Response{
			Message:    "Not a valid method",
			StatusCode: 405,
			Data:       "",
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the response for `Not a valid method`")
			return
		}
		return
	}

	order := new(shared.Order)
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		res := shared.Response{
			Message:    "Error while decoding the body",
			StatusCode: 500,
			Data:       "",
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the response for `All fields are mandatory`")
			return
		}
		return
	}
	fmt.Println(order)

	count := len(order.Products)
	if count == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := shared.Response{
			Message:    "No products added",
			StatusCode: 400,
			Data:       "",
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the response for `Invalid value / Missing fields`")
			return
		}
		return
	}

	if order.Status == "" || order.Payment.PaymentMethod == "" || order.Payment.TotalAmount == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := shared.Response{
			Message:    "Invalid value / Missing fields",
			StatusCode: 400,
			Data:       "",
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			log.Println("Failed to encode the response for `Invalid value / Missing fields`")
			return
		}
		return
	}

	services.OrderProduct(order)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Println("Failed to fetch the response for `handlers.AddProduct`")
		return
	}

}
