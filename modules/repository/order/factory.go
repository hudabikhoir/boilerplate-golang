package order

import (
	"insanulab/business/order"
	"insanulab/util"
)

//RepositoryFactory Will return business.order.Repository based on resource rest api
func RepositoryFactory(dbCon *util.DatabaseConnection) order.Repository {
	orderRepo := NewCouchDBRepository(dbCon.CouchDBClient)

	return orderRepo
}
