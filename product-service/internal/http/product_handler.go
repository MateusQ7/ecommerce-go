package http

import (
	"net/http"

	"github.com/MateusQ7/ecommerce-go/product-service/internal/domain"
	"github.com/MateusQ7/ecommerce-go/product-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *usecase.ProductService
}

func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p domain.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
