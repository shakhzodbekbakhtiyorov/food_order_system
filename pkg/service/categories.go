package service

import (
	"errors"
	"gogogo"
	"gogogo/pkg/repository"
)

type CategoryService struct {
	repository repository.Category
}

func NewCategoryService(repository repository.Category) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) Create(category gogogo.CreateCategoryInput) (int, error) {
	old_category, err := s.repository.GetCategoryByName(category.Name)
	if err != nil {
		return 0, err
	}
	if (gogogo.Category{}) != old_category {
		return 0, errors.New("category with this name already exists")
	}
	return s.repository.CreateCategory(category)
}

func (s *CategoryService) GetAll() ([]gogogo.Category, error) {
	return s.repository.GetAllCategories()
}

func (s *CategoryService) GetById(id int) (gogogo.Category, error) {
	return s.repository.GetCategoryById(id)
}

func (s *CategoryService) Delete(id int) error {
	return s.repository.DeleteCategory(id)
}

func (s *CategoryService) Update(id int, input gogogo.CreateCategoryInput) error {
	return s.repository.UpdateCategory(id, input)
}
