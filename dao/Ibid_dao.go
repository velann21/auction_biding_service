package dao

import (
	"context"
	"github.com/auction_biding/entities/requests"
)

type BidDao interface {
	userBidRecorder(userID string, data *requests.BidRequest)error;
	ProductBidRecorder(productID string, data *requests.BidRequest)error;
	TakeNewRecord(ctx context.Context,actionType string, data *requests.BidRequest);
    TaskExecutor();
    GetBidByUser(ctx context.Context, userID string)[]*requests.BidRequest
	GetBidByItem(ctx context.Context, productID string)[]*requests.BidRequest
    GetWinnerBidByItem(ctx context.Context, productID string) *requests.BidRequest
}
