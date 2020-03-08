package requests

import (
	"encoding/json"
	"github.com/auction_biding/helpers"
	"github.com/google/uuid"
	"io"
	"log"
	"strconv"
	"time"
)

type BidRequest struct {
	UserID    string  `json:"userID"`
	ProductID string  `json:"productID"`
	BidPrice  float64 `json:"bidPrice"`
	TimeStamp string  `json:"timeStamp"`
	UUID string
}

func (bidRequest *BidRequest)GetUserID()string{
     return bidRequest.UserID
}

func (bidRequest *BidRequest)GetProductID()string{
     return bidRequest.ProductID
}

func (bidRequest *BidRequest)GetBidPrice()float64{
	return bidRequest.BidPrice
}

func (bidRequest *BidRequest)GetTimeStamp()string{
	return bidRequest.TimeStamp
}

func (bidRequest *BidRequest) PopulateBidRequests(body io.ReadCloser) error {
	bidRequest.UUID =  uuid.New().String()
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&bidRequest)
	if err != nil {
		return helpers.NotValidRequestBody
	}
	return nil
}

func (bidRequest *BidRequest) ValidateBidRequest() error {

	if bidRequest.GetUserID() == "" || bidRequest.GetUserID() == " "{
		return helpers.NotValidRequestBody
	}

	if bidRequest.GetProductID() == "" || bidRequest.GetProductID() == " "{
		return helpers.NotValidRequestBody
	}

	if bidRequest.GetBidPrice() < 0{
		return helpers.NotValidRequestBody
	}

	_, err := time.Parse(time.RFC3339, bidRequest.GetTimeStamp())
	if err != nil {
		return helpers.NotValidRequestBody
	}

	return nil
}

func (bidRequest BidRequest) MarshalBidRequestObject(data BidRequest)([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}


type GetBidByUser struct {
	ID int `json:"id"`
}

func (getBidByUser *GetBidByUser)PopulateGetBidByUser(id string) error{
	intID, err := strconv.Atoi(id)
	if err != nil{
		return helpers.NotValidRequestBody
	}
	getBidByUser.ID = intID
	return nil
}

func (getBidByUser *GetBidByUser) ValidateGetBidByUser() error{
	if getBidByUser.ID < 0{
		return helpers.NotValidRequestBody
	}
	return nil
}

type GetBidByItem struct {
	ID int `json:"id"`
}

func (getBidByItem *GetBidByItem)PopulateGetBidByItem(id string) error{
	intID, err := strconv.Atoi(id)
	if err != nil{
		return helpers.NotValidRequestBody
	}
	getBidByItem.ID = intID
	return nil
}

func (getBidByItem *GetBidByItem) ValidateGetBidByItem() error{
	if getBidByItem.ID < 0{
		return helpers.NotValidRequestBody
	}
	return nil
}


type GetWinningBidByItem struct {
	ID int `json:"id"`
}

func (getWinningBidByItem *GetWinningBidByItem)PopulateGetWinningBidByItem(id string) error{
	intID, err := strconv.Atoi(id)
	if err != nil{
		return helpers.NotValidRequestBody
	}
	getWinningBidByItem.ID = intID
	return nil
}

func (getWinningBidByItem *GetWinningBidByItem) ValidateGetWinningBidByItem() error{
	if getWinningBidByItem.ID < 0{
		return helpers.NotValidRequestBody
	}
	return nil
}
