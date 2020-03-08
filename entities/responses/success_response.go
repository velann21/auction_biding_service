package responses

import (
	"encoding/json"
	"github.com/auction_biding/entities/requests"
	"log"
	"net/http"
)

// Response struct
type Response struct {
	Status string                   `json:"status"`
	Data   []map[string]interface{} `json:"data"`
	Meta   map[string]interface{}   `json:"meta,omitempty"`
}

func (entity *Response) GetBidByUserResponse(respData []*requests.BidRequest){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["UserBids"] = respData
	responseData = append(responseData, data)
	entity.Data = responseData
}

func (entity *Response) GetBidByItemResponse(respData []*requests.BidRequest){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["ProductBids"] = respData
	responseData = append(responseData, data)
	entity.Data = responseData
}

func (entity *Response) GetWinnerBidByItemResponse(respData *requests.BidRequest){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["WinnerBid"] = respData
	responseData = append(responseData, data)
	entity.Data = responseData
}

type products struct{
	ProductID string `json:"productID"`
	ProductName string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	Availability bool `json:"availability"`
	CanBid bool `json:"canBid"`
}
func (entity *Response) ListProducts(){

	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	products := []products{
		{"1","Macbook", "For coding", true, true},
		{"2","Philps TV", "For Watching Tv", true, true},
		{"3","Lenovo", "For coding", true, true},
		{"4","Mac Air", "For coding", true, true},
		{"5","PS4", "For gaming", true, true},
		{"6","Ps3", "For gaming", true, true},
		{"7","Fridge", "For Cooling", true, true},
		{"8","Apple phone", "For surfing", true, true},
		{"9","Chair", "For sitting", true, true},
		{"10","Table", "For eating", true, true},
		{"11","Shoes", "For wearing", true, true},
	}
	data["Products"] = products
	responseData = append(responseData, data)
	entity.Data = responseData

}

// SendResponse send http response
func (entity *Response) SendResponse(rw http.ResponseWriter, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		entity.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		entity.Status = http.StatusText(http.StatusAccepted)
	default:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	}

	// send response
	err := json.NewEncoder(rw).Encode(entity)
	if err != nil{
		log.Fatal("Something wrong")
	}
	return
}
