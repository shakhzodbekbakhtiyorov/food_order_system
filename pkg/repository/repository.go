package repository

import (
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gogogo.User) (int, error)
	GetUser(username, password string) (gogogo.User, error)
}

type Product interface {
	CreateProduct(product gogogo.CreateProductInput) (int, error)
	GetProductById(id int) (gogogo.Product, error)
	GetProductByName(name string) (gogogo.Product, error)
	GetAllProducts() ([]gogogo.Product, error)
	UpdateProduct(id int, input gogogo.UpdateProductInput) error
	DeleteProduct(id int) error
}

type Category interface {
	CreateCategory(category gogogo.CreateCategoryInput) (int, error)
	GetCategoryById(id int) (gogogo.Category, error)
	GetCategoryByName(name string) (gogogo.Category, error)
	GetAllCategories() ([]gogogo.Category, error)
	UpdateCategory(id int, input gogogo.CreateCategoryInput) error
	DeleteCategory(id int) error
}

type Order interface {
	CreateOrder() (int, error)
	GetOrderById(id int) (gogogo.OrderDB, error)
	GetAllOrders() ([]gogogo.Order, error)
	UpdateOrder(id int, input gogogo.UpdateOrderInput) error
	DeleteOrder(id int) error
}

type OrderItems interface {
	CreateOrderItems(order_id int, items []gogogo.CreateOrderItem) error
	CreateOrderItem(order_id int, item gogogo.CreateOrderItem) (int, error)
	GetOrderItemById(id int) (gogogo.OrderItem, error)
	GetOrderItemsByOrderId(order_id int) ([]gogogo.OrderItem, error)
	UpdateOrderItem(id int, input gogogo.UpdateOrderItem) error
	DeleteOrderItem(id int) error
}

type Payments interface {
	CreatePayment(input gogogo.CreatePaymentInput) (int, error)
	GetPaymentById(id int) (gogogo.Payment, error)
	GetPaymentByOrderId(order_id int) (gogogo.Payment, error)
	GetAllPayments() ([]gogogo.Payment, error)
	UpdatePayment(id int, input gogogo.UpdatePaymentInput) error
	DeletePayment(id int) error
}

type Repository struct {
	Authorization
	Product
	Category
	Order
	OrderItems
	Payments
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		Product:       NewProductRepository(db),
		Category:      NewCategoryRepository(db),
		Order:         NewOrderRepository(db),
		OrderItems:    NewOrderItemsRepository(db),
		Payments:      NewPaymentsRepository(db),
	}
}
