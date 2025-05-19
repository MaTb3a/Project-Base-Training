package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MaTb3aa/Project-Base-Training/config"
	"github.com/MaTb3aa/Project-Base-Training/docs"
	"github.com/MaTb3aa/Project-Base-Training/handlers"
	"github.com/MaTb3aa/Project-Base-Training/models"
	"github.com/MaTb3aa/Project-Base-Training/repository"
	"github.com/MaTb3aa/Project-Base-Training/routes"
	"github.com/MaTb3aa/Project-Base-Training/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/joho/godotenv"
	"wikidocify/internal/utils"
)

func waitForDependencies() {
	if err := utils.WaitForService("db", "5432", 30*time.Second); err != nil {
		log.Fatal("Database not ready:", err)
	}
	if err := utils.WaitForService("minio", "9000", 30*time.Second); err != nil {
		log.Fatal("MinIO not ready:", err)
	}
}

func connectDatabase(dsn string, maxAttempts int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, pingErr := db.DB()
			if pingErr == nil && sqlDB.Ping() == nil {
				log.Printf("✅ Connected to database on attempt %d", attempt)
				return db, nil
			}
		}
		log.Printf("⚠️  Attempt %d/%d failed: %v", attempt, maxAttempts, err)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("database connection failed: %w", err)
}

// @title Documents Service API
// @version 1.0
// @description API for managing documents.
// @host localhost:8080
// @BasePath /

func main() {
	waitForDependencies()

	cfg := config.LoadConfig()
	gin.SetMode(cfg.GinMode)

	db, err := connectDatabase(cfg.GetDSN(), 10, 2*time.Second)
	if err != nil {
		log.Fatal("❌ Could not connect to database:", err)
	}

	if err := db.AutoMigrate(&models.Document{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Database migration completed")

	repo := repository.NewGormRepository[models.Document](db)
	service := services.NewDocumentService(repo)
	handler := handlers.NewDocumentHandler(service)

	docs.SwaggerInfo.Host = cfg.SwaggerHost

	r := routes.SetupRouter(handler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("🚀 Starting server on :%s", cfg.APIPort)
	if err := r.Run(":" + cfg.APIPort); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
