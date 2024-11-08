package repository

import (
	"fmt"
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type OrderItemsRepository struct {
	db *sqlx.DB
}

func NewOrderItemsRepository(db *sqlx.DB) *OrderItemsRepository {
	return &OrderItemsRepository{db: db}
}

func (r *OrderItemsRepository) CreateOrderItems(order_id int, items []gogogo.CreateOrderItem) error {
	query := fmt.Sprintf("INSERT INTO %s (order_id, product_id, count, status) VALUES ($1, $2, $3, $4) RETURNING id", orderItemsTable)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, item := range items {
		_, err := stmt.Exec(order_id, item.ProductId, item.Count, 1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *OrderItemsRepository) CreateOrderItem(order_id int, item gogogo.CreateOrderItem) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (order_id, product_id, count, status) VALUES ($1, $2, $3, $4) RETURNING id", orderItemsTable)
	row := r.db.QueryRow(query, item.OrderId, item.ProductId, item.Count, 1)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderItemsRepository) GetOrderItemById(id int) (gogogo.OrderItem, error) {
	var orderItem gogogo.OrderItem
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", orderItemsTable)
	err := r.db.Get(&orderItem, query, id)
	return orderItem, err
}

func (r *OrderItemsRepository) GetOrderItemsByOrderId(order_id int) ([]gogogo.OrderItem, error) {
	var orderItems []gogogo.OrderItem
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", orderItemsTable)
	err := r.db.Select(&orderItems, query, order_id)
	return orderItems, err
}

func (r *OrderItemsRepository) UpdateOrderItem(id int, input gogogo.UpdateOrderItem) error {
	query := fmt.Sprintf("UPDATE status=$1, count=$2 SET %s WHERE id = $3",
		orderItemsTable)

	_, err := r.db.Exec(query, input.Status, input.Count, id)
	return err
}

func (r *OrderItemsRepository) DeleteOrderItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		orderItemsTable)
	_, err := r.db.Exec(query, id)
	return err
}
