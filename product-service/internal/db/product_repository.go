package db

import (
	"database/sql"
	"time"

	"github.com/MateusQ7/ecommerce-go/product-service/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{db: database}
}

func (r *ProductRepository) Save(product *domain.Product) error {
	query := `
		INSERT INTO products (name, description, price, stock_quantity, image_url, category_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.StockQuantity,
		product.ImageURL,
		product.CategoryID,
		time.Now(),
	).Scan(&product.ID, &product.CreatedAt)
}

func (r *ProductRepository) FindAll() ([]domain.Product, error) {
	query := `SELECT id, name, description, price, stock_quantity, image_url, category_id, created_at FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.StockQuantity,
			&p.ImageURL,
			&p.CategoryID,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
