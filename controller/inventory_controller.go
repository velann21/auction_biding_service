package controller

import (
	"github.com/auction_biding/entities/responses"
	"net/http"
)


func ListProducts(rw http.ResponseWriter, req *http.Request){
	resp := responses.Response{}
	resp.ListProducts()
	resp.SendResponse(rw, 200)
	return
}
