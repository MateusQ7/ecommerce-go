package db

import (
	"database/sql"

	"github.com/MateusQ7/ecommerce-go/product-service/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{db: database}
}

func (r *ProductRepository) Save(product *domain.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	return err
}

func (r *ProductRepository) FindAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var p domain.Product

		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
