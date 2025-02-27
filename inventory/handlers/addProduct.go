package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/likhithkp/ecommerce-order-managent-system/inventory/services"
	"github.com/likhithkp/ecommerce-order-managent-system/inventory/shared"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
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

	item := new(shared.Item)

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
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

	if item.Category == "" || item.Name == "" || item.Description == "" || item.Count < 0 || item.Price == 0 {
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

	res := services.AddProduct(item)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Println("Failed to fetch the response for `handlers.AddProduct`")
		return
	}

}
