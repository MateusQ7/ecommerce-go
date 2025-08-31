package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/MateusQ7/ecommerce-go/product-service/internal/db"
	"github.com/MateusQ7/ecommerce-go/product-service/internal/http"
	"github.com/MateusQ7/ecommerce-go/product-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSSLMode)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := db.NewProductRepository(dbConn)
	service := usecase.NewProductService(repo)
	handler := http.NewProductHandler(service)

	router := gin.Default()
	router.POST("/products", handler.CreateProduct)
	router.GET("/products", handler.ListProducts)

	log.Println("Product Server is running on :8080")
	router.Run(":8080")
}
