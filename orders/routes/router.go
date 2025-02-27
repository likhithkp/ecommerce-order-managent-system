package routes

import (
	"net/http"

	"github.com/likhithkp/ecommerce-order-managent-system/orders/handlers"
)

func InventoryRouter(mux *http.ServeMux) {
	mux.HandleFunc("/orderProduct", handlers.OrderProduct)
}
