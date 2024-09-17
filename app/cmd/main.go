package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/meter-peter/icsd-cloud-2024/docs"
	"github.com/meter-peter/icsd-cloud-2024/internal/models"
	"github.com/meter-peter/icsd-cloud-2024/internal/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title User Management API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8080
// @BasePath /
func main() {
	// Try to load .env file, if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	//If not found use the static alternatives
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "your_database_name")
	dbPort := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate the schema
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Add Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := getEnv("SERVER_PORT", "8080")
	r.Run(":" + port)
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
