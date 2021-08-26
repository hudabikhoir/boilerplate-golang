package modules

import (
	"insanulab/api"
	"insanulab/api/common"
	"insanulab/util"

	commodityCtrlV1 "insanulab/api/v1/commodity"
	commodityBussiness "insanulab/business/commodity"
	commodityRepo "insanulab/modules/repository/commodity"

	orderCtrlV1 "insanulab/api/v1/order"
	orderBussiness "insanulab/business/order"
	orderRepo "insanulab/modules/repository/order"

	echo "github.com/labstack/echo"
)

//SetErrorHandler - set error response
func SetErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// error message must be known RC value
		errResp := common.NewInternalServerErrorResponse()
		c.JSON(500, errResp)
	}
}

//RegisterController - register the controller
func RegisterController(dbCon *util.DatabaseConnection) api.Controller {

	//initiate commodity
	commodityPermitRepo := commodityRepo.RepositoryFactory()
	commodityPermitService := commodityBussiness.NewService(commodityPermitRepo)
	commodityPermitControllerV1 := commodityCtrlV1.NewController(commodityPermitService)

	//initiate order
	orderPermitRepo := orderRepo.RepositoryFactory(dbCon)
	orderPermitService := orderBussiness.NewService(orderPermitRepo)
	orderPermitControllerV1 := orderCtrlV1.NewController(orderPermitService)

	//lets put the controller together
	controllers := api.Controller{
		CommodityController: commodityPermitControllerV1,
		OrderController:     orderPermitControllerV1,
	}

	return controllers
}
