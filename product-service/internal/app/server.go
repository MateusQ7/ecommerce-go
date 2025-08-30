package app

import (
	"database/sql"
	"log"

	"github.com/MateusQ7/ecommerce-go/product-service/internal/db"
	"github.com/MateusQ7/ecommerce-go/product-service/internal/http"
	"github.com/MateusQ7/ecommerce-go/product-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	dbConection, err := sql.Open("postgres", "postgresql://root:admin@localhost:5432/products?sslmode=disable")

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := db.NewProductRepository(dbConection)
	service := usecase.NewProductService(repo)
	handler := http.NewProductHandler(service)

	router := gin.Default()

	router.POST("/products", handler.CreateProduct)
	router.GET("/products", handler.ListProducts)

	log.Println("Product Server is running on :8080")
	router.Run(":8080")
}
