package usecase

import (
	"github.com/MateusQ7/ecommerce-go/product-service/internal/domain"
)

type ProductRepository interface {
	Save(product *domain.Product) error
	FindAll() ([]domain.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (ps *ProductService) CreateProduct(product *domain.Product) error {
	return ps.repo.Save(product)
}

func (ps *ProductService) ListProducts() ([]domain.Product, error) {
	return ps.repo.FindAll()
}
