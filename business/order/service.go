package order

import "github.com/go-playground/validator/v10"

//Repository outgoing port for order
type Repository interface {
	//FetchOrders is a function to fetch orders from resource data
	FetchOrders() ([]Order, error)
}

//Service ingoing port for order
type Service interface {
	//GetOrders is a function to get order list service
	GetOrders() ([]Order, error)
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct order service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

func (s *service) GetOrders() ([]Order, error) {
	orders, err := s.repository.FetchOrders()
	if err != nil {
		return orders, err
	}
	return orders, nil
}
