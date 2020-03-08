package service

import (
	"context"
	"github.com/auction_biding/dao"
	dm "github.com/auction_biding/dao/dao_layer_dependency_manager"
	"github.com/auction_biding/entities/requests"
	"github.com/auction_biding/helpers"
)


type BidingServiceImpl struct {

}

/**
This UserBidRecorderService does have the business logics for User bids recorder
*/
func (bidingService BidingServiceImpl) UserBidRecorderService(ctx context.Context, request *requests.BidRequest) error {
	bidDao := dm.BidDaoDependencyManager(dm.BIDINGDAO)
	resp := dao.FindMaxBidInHeap(request.ProductID)
	if resp != nil{
		if resp.BidPrice >= request.BidPrice{
			return helpers.NotValidPrice
		}
	}
	go bidDao.TakeNewRecord(ctx, dao.USERBASEDBID, request)
	go bidDao.TakeNewRecord(ctx, dao.PRODUCTBASEDBID, request)
	return nil
}

/**
This GetBidsByItemservice does have the business logics for get the bids by item
*/
func (bidingService BidingServiceImpl)GetBidsByItemservice(ctx context.Context, productID string)([]*requests.BidRequest, error){
	bidDao := dm.BidDaoDependencyManager(dm.BIDINGDAO)
	resp := bidDao.GetBidByItem(ctx, productID)
	return resp, nil
}

/**
This GetWinnerBidService does have the business logics for get the bids winner by item
*/
func (bidingService BidingServiceImpl)GetWinnerBidService(ctx context.Context, productID string)(*requests.BidRequest, error){
	bidDao := dm.BidDaoDependencyManager(dm.BIDINGDAO)
	resp := bidDao.GetWinnerBidByItem(ctx, productID)
	return resp, nil
}

/**
This GetBidsByUsersService does have the business logics for get the bids by user
*/
func (bidingService BidingServiceImpl)GetBidsByUsersService(ctx context.Context,userID string)([]*requests.BidRequest, error){
	bidDao := dm.BidDaoDependencyManager(dm.BIDINGDAO)
	resp := bidDao.GetBidByUser(ctx, userID)
	return resp, nil
}


