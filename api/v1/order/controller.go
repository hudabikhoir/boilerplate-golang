package order

import (
	"insanulab/api/common"
	"insanulab/api/v1/order/response"
	orderBusiness "insanulab/business/order"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get order API controller
type Controller struct {
	service   orderBusiness.Service
	validator *v10.Validate
}

//NewController Construct order API controller
func NewController(service orderBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//GetAllOrders Get order by ID echo handler
func (controller *Controller) GetAllOrders(c echo.Context) error {
	commodities, err := controller.service.GetOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetOrdersResponse(commodities)
	return c.JSON(http.StatusOK, response)
}
