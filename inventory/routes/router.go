package routes

import (
	"net/http"

	"github.com/likhithkp/ecommerce-order-managent-system/inventory/handlers"
)

func InventoryRouter(mux *http.ServeMux) {
	mux.HandleFunc("/addProduct", handlers.AddProduct)
}
