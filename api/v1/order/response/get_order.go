package response

import (
	"insanulab/business/order"
)

//GetOrderByIDResponse Get order by ID response payload
type GetOrderByIDResponse struct {
	UUID    string `json:"uuid"`
	OrderID string `json:"order_id"`
	NoPKS   string `json:"no_pks"`
}

//GetCommoditieResponse Get commoditie by tag response payload
type GetOrdersResponse struct {
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Data    []*GetOrderByIDResponse `json:"data"`
}

//NewGetOrderByIDResponse construct GetOrderByIDResponse
func NewGetOrderByIDResponse(order order.Order) *GetOrderByIDResponse {
	var orderResponse GetOrderByIDResponse
	orderResponse.UUID = order.ID
	orderResponse.OrderID = order.OrderID
	orderResponse.NoPKS = order.NoPKS

	return &orderResponse
}

//NewGetCommoditieResponse construct GetCommoditieResponse
func NewGetOrdersResponse(commodities []order.Order) *GetOrdersResponse {
	var commoditiesResponses []*GetOrderByIDResponse
	commoditiesResponses = make([]*GetOrderByIDResponse, 0)

	for _, commodities := range commodities {
		commoditiesResponses = append(commoditiesResponses, NewGetOrderByIDResponse(commodities))
	}

	return &GetOrdersResponse{
		"00",
		"Success",
		commoditiesResponses,
	}
}
