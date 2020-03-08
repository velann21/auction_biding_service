package main

import (
	dm "github.com/auction_biding/dao/dao_layer_dependency_manager"
	"github.com/auction_biding/routes"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// This is the Bootsrap class..
func main() {

	daoObject := dm.BidDaoDependencyManager(dm.BIDINGDAO)
	go daoObject.TaskExecutor()
	// Routing logic
	r := mux.NewRouter().StrictSlash(false)
	mainRoutes := r.PathPrefix("/api/v1").Subrouter()
	routes.BidingRoutes(mainRoutes)
	routes.InventoryRoutes(mainRoutes)
	logrus.Info("Server Started")
	//Bootingup the server
	err := http.ListenAndServe(":8080", r)
	if err != nil{
		logrus.Fatal("Server start failed")
	}
}