package service

import (
	"gogogo"
	"gogogo/pkg/repository"
)

type PaymentService struct {
	repository repository.Payments
}

func NewPaymentService(repository repository.Payments) *PaymentService {
	return &PaymentService{
		repository: repository,
	}
}

func (s *PaymentService) Create(payment gogogo.CreatePaymentInput) (int, error) {
	return s.repository.CreatePayment(payment)
}

func (s *PaymentService) GetAll() ([]gogogo.Payment, error) {
	return s.repository.GetAllPayments()
}

func (s *PaymentService) GetById(id int) (gogogo.Payment, error) {
	return s.repository.GetPaymentById(id)
}

func (s *PaymentService) GetByOrderId(order_id int) (gogogo.Payment, error) {
	return s.repository.GetPaymentByOrderId(order_id)
}

func (s *PaymentService) Delete(id int) error {
	return s.repository.DeletePayment(id)
}

func (s *PaymentService) Update(id int, input gogogo.UpdatePaymentInput) error {
	return s.repository.UpdatePayment(id, input)
}
