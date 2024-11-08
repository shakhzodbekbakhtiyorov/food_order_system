package service

import (
	"errors"
	"fmt"
	"gogogo"
	"gogogo/pkg/repository"
)

type ProductService struct {
	repository repository.Product
}

func NewProductService(repository repository.Product) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Create(product gogogo.CreateProductInput) (int, error) {
	fmt.Println("Product: ", product)
	old_product, err := s.repository.GetProductByName(product.Name)
	if err != nil {
		return 0, err
	}
	fmt.Println("SEDONF")
	if (gogogo.Product{}) != old_product {
		return 0, errors.New("product with this name already exists")
	}
	return s.repository.CreateProduct(product)
}

func (s *ProductService) GetAll() ([]gogogo.Product, error) {
	return s.repository.GetAllProducts()
}

func (s *ProductService) GetById(id int) (gogogo.Product, error) {
	return s.repository.GetProductById(id)
}

func (s *ProductService) Delete(id int) error {
	return s.repository.DeleteProduct(id)
}

func (s *ProductService) Update(id int, input gogogo.UpdateProductInput) error {
	return s.repository.UpdateProduct(id, input)
}
