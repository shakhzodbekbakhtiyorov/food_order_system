package service

import (
	"gogogo"
	"gogogo/pkg/repository"
)

type OrderService struct {
	OrderRepo     repository.Order
	OrderItemRepo repository.OrderItems
}

func NewOrderService(orderRepo repository.Order, orderItemRepo repository.OrderItems) *OrderService {
	return &OrderService{
		OrderRepo:     orderRepo,
		OrderItemRepo: orderItemRepo,
	}
}

func (s *OrderService) Create(order_items gogogo.CreateOrderInput) (int, error) {
	order_id, err := s.OrderRepo.CreateOrder()
	if err != nil {
		return 0, err
	}
	err = s.OrderItemRepo.CreateOrderItems(order_id, order_items.Items)
	if err != nil {
		return 0, err
	}
	return order_id, nil
}

func (s *OrderService) GetAll() ([]gogogo.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

func (s *OrderService) GetById(id int) (gogogo.Order, error) {
	var order gogogo.Order
	orderdb, err := s.OrderRepo.GetOrderById(id)
	if err != nil {
		return order, err
	}

	orderItems, err := s.OrderItemRepo.GetOrderItemsByOrderId(id)
	if err != nil {
		return order, err
	}

	order.Order = orderdb
	order.OrderItems = orderItems

	return order, nil
}

func (s *OrderService) Delete(id int) error {
	return s.OrderRepo.DeleteOrder(id)
}

func (s *OrderService) Update(id int, input gogogo.UpdateOrderInput) error {
	return s.OrderRepo.UpdateOrder(id, input)
}
