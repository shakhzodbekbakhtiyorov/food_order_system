package service

import (
	"gogogo"
	"gogogo/pkg/repository"
)

type OrderItemService struct {
	OrderItemRepo repository.OrderItems
}

func NewOrderItemService(orderItemRepo repository.OrderItems) *OrderItemService {
	return &OrderItemService{
		OrderItemRepo: orderItemRepo,
	}
}

func (s *OrderItemService) CreateOrderItem(input gogogo.CreateOrderItem) (int, error) {
	return s.OrderItemRepo.CreateOrderItem(input.OrderId, input)
}

func (s *OrderItemService) EditOrderItem(id int, input gogogo.UpdateOrderItem) error {
	return s.OrderItemRepo.UpdateOrderItem(id, input)
}

func (s *OrderItemService) DeleteOrderItem(id int) error {
	return s.OrderItemRepo.DeleteOrderItem(id)
}
