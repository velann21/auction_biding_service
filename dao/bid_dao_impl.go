package dao

import (
	"context"
	"github.com/auction_biding/entities/requests"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"sync"
)

type BidDaoImpl struct {

}

const (
	USERBASEDBID = "addUserBasedBid"
	PRODUCTBASEDBID = "addProductBasedBid"
)

// Logfile to make data Durability
const (
	USERDATALOG = "./logfiles/user_data.log"
	PRODUCTDATALOG = "./logfiles/product_data.log"
	LASTPROCESSEDUSERDATA = "./logfiles/last_processed_product_data.log"
	LASTPROCESSEDPRODUCTLOG = "./logfiles/last_processed_user_data.log"
)

//Read write mutex for user map
var usersyncMutex sync.RWMutex
//Read write mutex for product map
var productsyncMutex sync.RWMutex

//to maintain the record of users based and product based separately
var userBids = make(map[string][]*requests.BidRequest,0)
var productBids = make(map[string][]*requests.BidRequest,0)

//To maintain the optimistic lock
var userBidsVersionTracker = make(map[string]float32,0)
var productBidsVersionTracker = make(map[string]float32,0)

//Buffered channel for user add and product add
var addUserBidRecordChan =  make(chan *requests.BidRequest, 100)
var addProductBidRecordChan =  make(chan *requests.BidRequest, 100)


func (bidDao BidDaoImpl) userBidRecorder(userID string, data *requests.BidRequest)error{
	logrus.Info("userBidRecorder Started")
	userBid := userBids[userID]
	userVersion := userBidsVersionTracker[userID]
	//If record for the user is first it creates a new entry
	if userBid == nil && userVersion == 0{
		usersyncMutex.RLock()
		if userBid == nil && userVersion == 0{
			logrus.WithField("ID",userID).Info("No user found creating new entry")
			bidsArray := make([]*requests.BidRequest, 0)
			bidsArray = append(bidsArray, data)
			reReadVersion := userBidsVersionTracker[userID]
			if reReadVersion == 0{
				userBids[userID] = bidsArray
				userBidsVersionTracker[userID] = 1.0
			}else{
				userBid := userBids[userID]
				userBid = append(userBid, data)
				userBids[userID] = userBid
				userVersion := userBidsVersionTracker[userID]
				userBidsVersionTracker[userID] = userVersion+1
			}

		}
		usersyncMutex.RUnlock()
	}else{
		// Else it append the data
		logrus.WithField("ID",userID).Info("User found creating appending the record to slice")
		userBid = append(userBid, data)
		userBids[userID] = userBid
		usersyncMutex.Lock()
		userVersion := userBidsVersionTracker[userID]
		userBidsVersionTracker[userID] = userVersion+1
		usersyncMutex.Unlock()
	}
	logrus.Info("userBidRecorder Done")
	return nil
}


/**
This GetWinnerBidService does have the business logics for get the bids winner by item
*/
func (bidDao BidDaoImpl) ProductBidRecorder(productID string, data *requests.BidRequest) error{
	logrus.Info("productBidRecorder Started")
	productsBid := productBids[productID]
	productVersion := productBidsVersionTracker[productID]

	//If record for the product is first it creates a new entry and make Heapify the slice
	if productsBid == nil && productVersion == 0{
		productsyncMutex.RLock()
		if productsBid == nil && productVersion == 0 {
			logrus.WithField("ID", productID).Info("No product found creating new entry")
			productBidsArray := make([]*requests.BidRequest, 1)
			productBidsArray = append(productBidsArray, data)
			productBids[productID] = productBidsArray
			reReadVersion := productBidsVersionTracker[productID]
			if reReadVersion == 0{
				productBidsVersionTracker[productID] = 1.0
			}else{
				productsBid := productBids[productID]
				productsBid = append(productsBid, data)
				productBids[productID] = productsBid
				productVersion := productBidsVersionTracker[productID]
				productBidsVersionTracker[productID] = productVersion+1
			}
		}
		productsyncMutex.RUnlock()
	}else{
		// Else it append the data and Heapify the slice to get the max bid in O(1)
		logrus.WithField("ID",productID).Info("Product found creating appending to slice")
		productsBid = append(productsBid, data)
		productBids[productID] = productsBid
		productsyncMutex.Lock()
		// This will make the heap to get the record in O(1)
		CreateHeap(productBids[productID])
		productVersion := productBidsVersionTracker[productID]
		productBidsVersionTracker[productID] = productVersion+1
		productsyncMutex.Unlock()
	}
	logrus.Info("productBidRecorder Done")
	return nil
}


/**
This GetBidByUser
*/
func (bidDao BidDaoImpl) GetBidByUser(ctx context.Context, userID string)[]*requests.BidRequest{
	userBid := userBids[userID]
	return userBid
}


/**
This GetBidByItem
*/
func (bidDao BidDaoImpl) GetBidByItem(ctx context.Context, productID string) []*requests.BidRequest{
	productsBID := productBids[productID]
	return productsBID
}


/**
This GetWinnerBidByItem
*/
func (bidDao BidDaoImpl) GetWinnerBidByItem(ctx context.Context, productID string) *requests.BidRequest{
	productsBID := productBids[productID]
	if len(productsBID) >= 2{
		return productsBID[1]
	}
	return nil
}

//Command Pattern
func (bidDao BidDaoImpl) TakeNewRecord(ctx context.Context, actionType string, data *requests.BidRequest){

	switch actionType {
	case USERBASEDBID:
		logrus.WithField("ID", data.UserID).Info(USERBASEDBID)
		WriteToLogFile(USERDATALOG, data.UUID,data.UserID, data)
        addUserBidRecordChan <- data

	case PRODUCTBASEDBID:
		logrus.WithField("ID", data.UserID).Info(PRODUCTBASEDBID)
		WriteToLogFile(PRODUCTDATALOG, data.UUID,data.UserID, data)
		addProductBidRecordChan <- data

	default:
		log.Print("No record Found")

	}
}

//Command Pattern Executor
func (bidDao BidDaoImpl) TaskExecutor(){
	for {
		select {
		case task := <-addUserBidRecordChan:
            go func(){
				err := bidDao.userBidRecorder(task.GetUserID(), task)
				if err != nil {
					log.Print("Error While adding the Bid For User bases: " + task.GetUserID())
				}
				WriteToLogFile(LASTPROCESSEDUSERDATA, task.UUID, task.UserID, task)
			}()

		case task := <-addProductBidRecordChan:
			go func() {
				err := bidDao.ProductBidRecorder(task.GetProductID(), task)
				if err != nil {
					log.Print("Error While adding the Bid For Product bases: " + task.GetProductID())
				}
				WriteToLogFile(LASTPROCESSEDPRODUCTLOG, task.UUID, task.ProductID, task)
			}()
		}
	}
}

func WriteToLogFile(fileName string, uuid string, id string, req *requests.BidRequest){
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, uuid+",", log.LstdFlags)
	logger.Println(id, req)
}


/**
This CreateHeap one by one for the slice
*/
func CreateHeap(array []*requests.BidRequest){
	for i :=2; i<=len(array)-1 ;i++{
		Insert(array, i)
	}

}

/**
This make the Heap form of an array rep
*/
func Insert(array []*requests.BidRequest, lastStoredOffset int)*requests.BidRequest{
	i := lastStoredOffset
	temp := array[lastStoredOffset]
	for i>1 && temp.BidPrice>array[i/2].BidPrice{
		array[i] = array[i/2]
		i = i/2
	}
	array[i] = temp
	return array[1]
}


/**
This FindMaxBidInHeap to find the max value in heap
*/
func FindMaxBidInHeap(productID string) *requests.BidRequest {
	productsBID := productBids[productID]
	if len(productsBID) >= 2{
		return productsBID[1]
	}
	return nil
}
