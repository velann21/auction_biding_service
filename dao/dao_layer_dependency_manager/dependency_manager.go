package dao_layer_dependency_manager

import "github.com/auction_biding/dao"

const(
	BIDINGDAO = "BidingDao"
)

func BidDaoDependencyManager(objectType string)dao.BidDao{

	if objectType == BIDINGDAO{
		return dao.BidDaoImpl{}
	}
	return nil
}

