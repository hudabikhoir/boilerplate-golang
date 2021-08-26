package api

import (
	"insanulab/api/middleware"
	"insanulab/api/v1/commodity"
	"insanulab/api/v1/order"
	serviceAuth "insanulab/business/auth"

	"github.com/labstack/echo"
)

// Controller to define controller that we use
type Controller struct {
	CommodityController *commodity.Controller
	OrderController     *order.Controller
}

//RegisterPath V1 API path
func RegisterPath(e *echo.Echo, ctrl Controller, auth serviceAuth.Service) {
	//commodity
	commodityV1 := e.Group("/v1/commodities")
	commodityV1.Use(middleware.JWTMiddleware(auth))
	commodityV1.GET("", ctrl.CommodityController.GetAllCommodities)
	commodityV1.GET("/report", ctrl.CommodityController.GetReportCommodities)

	orderV1 := e.Group("/v1/orders")
	orderV1.GET("", ctrl.OrderController.GetAllOrders)
	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
