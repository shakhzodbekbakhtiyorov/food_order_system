package repository

import (
	"database/sql"
	"fmt"
	"gogogo"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product gogogo.CreateProductInput) (int, error) {
	fmt.Println("PRODYCT: ", product)
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, price, category_id, image) values ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Price, product.CategoryId, product.Image)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) GetProductById(id int) (gogogo.Product, error) {
	var product gogogo.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, id)
	return product, err
}

func (r *ProductRepository) GetProductByName(name string) (gogogo.Product, error) {
	var product gogogo.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", productsTable)
	err := r.db.Get(&product, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return gogogo.Product{}, nil
		}
		return gogogo.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) GetAllProducts() ([]gogogo.Product, error) {
	var products []gogogo.Product

	query := fmt.Sprintf("SELECT * FROM %s", productsTable)

	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductRepository) UpdateProduct(id int, input gogogo.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.CategoryId != nil {
		setValues = append(setValues, fmt.Sprintf("category_id=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		productsTable, setQuery, argId)

	fmt.Println(query)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		productsTable)
	_, err := r.db.Exec(query, id)
	return err
}
