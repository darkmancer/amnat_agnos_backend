package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strong_password_recommendation/infrastructure"
	"strong_password_recommendation/internal/core/repository"
	"strong_password_recommendation/internal/core/service"
	"strong_password_recommendation/internal/handler/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	fmt.Println("Current directory:", dir)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	db, err := infrastructure.NewPostgresDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	logRepo := repository.NewLogRepository(db)
	passwordService := service.NewPasswordService(logRepo)

	router := gin.Default()
	api.SetupRouter(router, passwordService)

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
