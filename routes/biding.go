package routes

import (
	 "github.com/gorilla/mux"
	 "github.com/auction_biding/controller"
	 )

func BidingRoutes(route *mux.Router){
	route.HandleFunc("/biding/usersBid", controller.UserBidRecorderController).Methods("POST")
	route.HandleFunc("/biding/getBidByItem", controller.GetBidsByItemController).Methods("GET")
	route.HandleFunc("/biding/getBidByUser", controller.GetBidsByUserController).Methods("GET")
	route.HandleFunc("/biding/getWinningBid", controller.GetWinnerBidController).Methods("GET")
}
