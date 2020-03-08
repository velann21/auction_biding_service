package controller

import (
	"github.com/auction_biding/entities/requests"
	"github.com/auction_biding/entities/responses"
	dm "github.com/auction_biding/service/service_layer_dependency_manager"
	"github.com/sirupsen/logrus"
	"net/http"
)

func UserBidRecorderController(rw http.ResponseWriter, req *http.Request){
	logrus.Info("UserBidRecorderController started")
	bidRequest := requests.BidRequest{}
	bidService := dm.BidServiceDependencyManager(dm.BIDINGSERVICE)
    successResp := responses.Response{}
	err := bidRequest.PopulateBidRequests(req.Body)
	if err != nil{
		logrus.WithError(err).Error("Error occured in PopulateBidRequests()")
		responses.HandleError(rw, err)
		return
	}
	err = bidRequest.ValidateBidRequest()
	if err != nil{
		logrus.WithError(err).Error("Error occured in ValidateBidRequest()")
		responses.HandleError(rw, err)
		return
	}
	err = bidService.UserBidRecorderService(req.Context(), &bidRequest)
	if err != nil{
		logrus.WithError(err).Error("Error occured in UserBidRecorderService()")
		responses.HandleError(rw, err)
		return
	}
	successResp.SendResponse(rw, http.StatusOK)
	logrus.Info("UserBidRecorderController done")
	return
}

func GetBidsByItemController(rw http.ResponseWriter, req *http.Request){
	logrus.Info("GetBidsByItemController Started")
	successResp := responses.Response{}
	request := requests.GetBidByItem{}
	id := req.URL.Query()["product_id"][0]
	err := request.PopulateGetBidByItem(id)
	if err != nil{
		logrus.WithError(err).Error("Error while PopulateGetBidByItem()")
		responses.HandleError(rw, err)
		return
	}
	err = request.ValidateGetBidByItem()
	if err != nil{
		logrus.WithError(err).Error("Error while ValidateGetBidByItem()")
		responses.HandleError(rw, err)
		return
	}
	bidService := dm.BidServiceDependencyManager(dm.BIDINGSERVICE)
	resp, err := bidService.GetBidsByItemservice(req.Context(),id)
	if err != nil{
		logrus.WithError(err).Error("Error while GetBidsByItemController()")
		responses.HandleError(rw, err)
		return
	}
	successResp.GetBidByItemResponse(resp)
	successResp.SendResponse(rw, http.StatusOK)
	logrus.Info("GetBidsByItemController done")
	return

}

func GetWinnerBidController(rw http.ResponseWriter, req *http.Request){
	logrus.Info("GetWinnerBidController Started")
	successResp := responses.Response{}
	request := requests.GetWinningBidByItem{}
	id := req.URL.Query()["product_id"][0]
	err := request.PopulateGetWinningBidByItem(id)
	if err != nil{
		logrus.WithError(err).Error("Error while PopulateGetWinningBidByItem()")
		responses.HandleError(rw, err)
		return
	}
	err = request.PopulateGetWinningBidByItem(id)
	if err != nil{
		logrus.WithError(err).Error("Error while PopulateGetWinningBidByItem()")
		responses.HandleError(rw, err)
		return
	}
	bidService := dm.BidServiceDependencyManager(dm.BIDINGSERVICE)
	resp, err := bidService.GetWinnerBidService(req.Context(), id)
	if err != nil{
		logrus.WithError(err).Error("Error while GetWinnerBidService()")
		responses.HandleError(rw, err)
		return
	}
	successResp.GetWinnerBidByItemResponse(resp)
	successResp.SendResponse(rw, http.StatusOK)
	logrus.Info("GetWinnerBidController done")
	return

}

func GetBidsByUserController(rw http.ResponseWriter, req *http.Request){
	logrus.Info("GetBidsByUserController Started")
	successResp := responses.Response{}
	request := requests.GetBidByUser{}
	id := req.URL.Query()["user_id"][0]
	err := request.PopulateGetBidByUser(id)
	if err != nil{
		logrus.WithError(err).Error("Error while PopulateGetBidByUser()")
		responses.HandleError(rw, err)
		return
	}
	err = request.ValidateGetBidByUser()
	if err != nil{
		logrus.WithError(err).Error("Error while ValidateGetBidByUser()")
		responses.HandleError(rw, err)
		return
	}
	bidService := dm.BidServiceDependencyManager(dm.BIDINGSERVICE)
	resp, err := bidService.GetBidsByUsersService(req.Context(), id)
	if err != nil{
		logrus.WithError(err).Error("Error while GetBidsByUsers()")
		responses.HandleError(rw, err)
		return
	}
	successResp.GetBidByUserResponse(resp)
	successResp.SendResponse(rw, http.StatusOK)
	logrus.Info("GetBidsByUserController done")
	return
}