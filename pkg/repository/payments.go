package repository

import (
	"fmt"
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type PaymentsRepository struct {
	db *sqlx.DB
}

func NewPaymentsRepository(db *sqlx.DB) *PaymentsRepository {
	return &PaymentsRepository{db: db}
}

func (r *PaymentsRepository) CreatePayment(input gogogo.CreatePaymentInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (order_id, payment_type, amount, change) values ($1, $2, $3, $4) RETURNING id", paymentsTable)
	row := r.db.QueryRow(query, input.OrderId, input.PaymentType, input.Amount, input.Change)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PaymentsRepository) GetPaymentById(id int) (gogogo.Payment, error) {
	var payment gogogo.Payment
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", paymentsTable)
	err := r.db.Get(&payment, query, id)
	return payment, err
}

func (r *PaymentsRepository) GetPaymentByOrderId(order_id int) (gogogo.Payment, error) {
	var payment gogogo.Payment
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", paymentsTable)
	err := r.db.Get(&payment, query, order_id)
	return payment, err
}

func (r *PaymentsRepository) GetAllPayments() ([]gogogo.Payment, error) {
	var payments []gogogo.Payment

	query := fmt.Sprintf("SELECT * FROM %s", paymentsTable)

	err := r.db.Select(&payments, query)
	return payments, err
}

func (r *PaymentsRepository) UpdatePayment(id int, input gogogo.UpdatePaymentInput) error {
	query := fmt.Sprintf("UPDATE payment_type=$1, amount=$2, change=$3 SET %s WHERE id = $4",
		paymentsTable)

	_, err := r.db.Exec(query, input.PaymentType, input.Amount, input.Change, id)
	return err
}

func (r *PaymentsRepository) DeletePayment(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		paymentsTable)
	_, err := r.db.Exec(query, id)
	return err
}
