package domain

import "time"

type Product struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	ImageURL      string    `json:"image_url"`
	CategoryID    int64     `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
}
