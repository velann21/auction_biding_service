package service_layer_dependency_manager

import "github.com/auction_biding/service"


const(
	BIDINGSERVICE = "BidingService"
)

func BidServiceDependencyManager(objectType string)service.Bid{
	if objectType == BIDINGSERVICE{
		return service.BidingServiceImpl{}
	}
	return nil
}

