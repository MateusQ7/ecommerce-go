package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MateusQ7/ecommerce-go/auth-service/internal/db"
	"github.com/MateusQ7/ecommerce-go/auth-service/internal/http"
	"github.com/MateusQ7/ecommerce-go/auth-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func StartServe() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSSLMode)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := db.NewUserRepository(dbConn)
	service := usecase.NewUserService(repo)
	handler := http.NewUserHandler(service)

	router := gin.Default()
	router.POST("/create-user", handler.CreateNewUser)
	router.GET("/users", handler.ListUsers)

	log.Println("Auth Server is running on: 8081")
	router.Run(":8081")
}
