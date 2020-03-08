package routes

import (
	"github.com/auction_biding/controller"
	"github.com/gorilla/mux"
)

func InventoryRoutes(route *mux.Router) {
	route.HandleFunc("/inventory/products", controller.ListProducts).Methods("GET")
}

