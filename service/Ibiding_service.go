package service

import (
    "context"
    "github.com/auction_biding/entities/requests"
)

type Bid interface {
    UserBidRecorderService(ctx context.Context, request *requests.BidRequest) error;
    GetBidsByItemservice(ctx context.Context, productID string) ([]*requests.BidRequest,error);
    GetWinnerBidService(ctx context.Context, productID string) (*requests.BidRequest,error);
    GetBidsByUsersService(ctx context.Context, userID string) ([]*requests.BidRequest,error);
}
