package repository

import (
	"database/sql"
	"fmt"
	"gogogo"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category gogogo.CreateCategoryInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, image) values ($1, $2) RETURNING id", categoriesTable)
	fmt.Println("category: ", query)
	row := r.db.QueryRow(query, category.Name, category.Image)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CategoryRepository) GetCategoryById(id int) (gogogo.Category, error) {
	var category gogogo.Category
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", categoriesTable)
	err := r.db.Get(&category, query, id)
	return category, err
}

func (r *CategoryRepository) GetCategoryByName(name string) (gogogo.Category, error) {
	var category gogogo.Category
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", categoriesTable)
	err := r.db.Get(&category, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return gogogo.Category{}, nil
		}
		return gogogo.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) GetAllCategories() ([]gogogo.Category, error) {
	var categories []gogogo.Category

	query := fmt.Sprintf("SELECT * FROM %s", categoriesTable)

	err := r.db.Select(&categories, query)
	return categories, err
}

func (r *CategoryRepository) UpdateCategory(id int, input gogogo.CreateCategoryInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, image=$2 WHERE id = $3",
		categoriesTable)

	_, err := r.db.Exec(query, input.Name, input.Image, id)
	return err
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		categoriesTable)
	_, err := r.db.Exec(query, id)
	return err
}
