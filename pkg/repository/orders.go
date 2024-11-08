package repository

import (
	"fmt"
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder() (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (order_time, price, category_id) values ($1, $2, $3) RETURNING id", ordersTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderRepository) GetOrderById(id int) (gogogo.OrderDB, error) {
	var order gogogo.OrderDB
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", ordersTable)
	err := r.db.Get(&order, query, id)
	return order, err
}

func (r *OrderRepository) GetorderByOrderId(order_id int) (gogogo.OrderDB, error) {
	var order gogogo.OrderDB
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", ordersTable)
	err := r.db.Get(&order, query, order_id)
	return order, err
}

func (r *OrderRepository) GetAllOrders() ([]gogogo.Order, error) {
	var ordersMap = make(map[int]gogogo.Order)
	var orders []gogogo.Order

	query := fmt.Sprintf("SELECT * FROM %s o LEFT JOIN %s oi ON o.id = oi.order_id", ordersTable, orderItemsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderDB gogogo.OrderDB
		var orderItem gogogo.OrderItem
		err := rows.Scan(
			&orderDB.Id,
			&orderDB.OrderTime,
			&orderDB.Status,
			&orderDB.OrderNumber,
			&orderDB.Amount,
			&orderItem.Id,
			&orderItem.Count,
			&orderItem.ProductId,
			&orderItem.OrderId,
			&orderItem.Status,
		)
		if err != nil {
			return nil, err
		}

		if _, ok := ordersMap[orderDB.Id]; !ok {
			ordersMap[orderDB.Id] = gogogo.Order{
				Order:      orderDB,
				OrderItems: []gogogo.OrderItem{},
			}
		}

		existingOrder := ordersMap[orderDB.Id]
		existingOrder.OrderItems = append(existingOrder.OrderItems, orderItem)
		ordersMap[orderDB.Id] = existingOrder
	}

	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) UpdateOrder(id int, input gogogo.UpdateOrderInput) error {
	query := fmt.Sprintf("UPDATE status=$1 SET %s WHERE id = $2",
		ordersTable)

	_, err := r.db.Exec(query, input.Status, id)
	return err
}

func (r *OrderRepository) DeleteOrder(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		ordersTable)
	_, err := r.db.Exec(query, id)
	return err
}
