package service

import (
	"gogogo"
	"gogogo/pkg/repository"
)

type Auhtorization interface {
	CreateUser(user gogogo.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Products interface {
	Create(product gogogo.CreateProductInput) (int, error)
	GetAll() ([]gogogo.Product, error)
	GetById(id int) (gogogo.Product, error)
	Delete(id int) error
	Update(id int, input gogogo.UpdateProductInput) error
}

type Categories interface {
	Create(product gogogo.CreateCategoryInput) (int, error)
	GetAll() ([]gogogo.Category, error)
	GetById(id int) (gogogo.Category, error)
	Delete(id int) error
	Update(id int, input gogogo.CreateCategoryInput) error
}

type Orders interface {
	Create(order_items gogogo.CreateOrderInput) (int, error)
	GetAll() ([]gogogo.Order, error)
	GetById(id int) (gogogo.Order, error)
	Delete(id int) error
	Update(id int, input gogogo.UpdateOrderInput) error
}

type OrderItems interface {
	CreateOrderItem(input gogogo.CreateOrderItem) (int, error)
	EditOrderItem(id int, input gogogo.UpdateOrderItem) error
	DeleteOrderItem(id int) error
}

type Payments interface {
	Create(payment gogogo.CreatePaymentInput) (int, error)
	GetAll() ([]gogogo.Payment, error)
	GetById(id int) (gogogo.Payment, error)
	GetByOrderId(order_id int) (gogogo.Payment, error)
	Delete(id int) error
	Update(id int, input gogogo.UpdatePaymentInput) error
}

type Service struct {
	Auhtorization
	Products
	Categories
	Orders
	OrderItems
	Payments
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auhtorization: NewAuthService(repository.Authorization),
		Products:      NewProductService(repository.Product),
		Categories:    NewCategoryService(repository.Category),
		Orders:        NewOrderService(repository.Order, repository.OrderItems),
		OrderItems:    NewOrderItemService(repository.OrderItems),
		Payments:      NewPaymentService(repository.Payments),
	}
}
