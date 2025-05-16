package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MaTb3aa/Project-Base-Training/docs"
	_ "github.com/MaTb3aa/Project-Base-Training/docs"
	Repositories "github.com/MaTb3aa/Project-Base-Training/repository"
	"github.com/MaTb3aa/Project-Base-Training/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MaTb3aa/Project-Base-Training/config"
	Handlers "github.com/MaTb3aa/Project-Base-Training/handlers"
	"github.com/MaTb3aa/Project-Base-Training/models"
	Services "github.com/MaTb3aa/Project-Base-Training/services"
	"github.com/gin-gonic/gin"
)

// connectDatabase attempts to open a GORM connection and ping the DB.
// It retries up to maxAttempts times, waiting delay between tries.
func connectDatabase(dsn string, maxAttempts int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, pingErr := db.DB()
			if pingErr == nil {
				if pingErr = sqlDB.Ping(); pingErr == nil {
					log.Printf("✅ Database connected on attempt %d", attempt)
					// Run migration AFTER DB is ready
					if migrateErr := db.AutoMigrate(&models.DocumentFromOrm{}); migrateErr != nil {
						return nil, fmt.Errorf("Migration failed: %w", migrateErr)
					}
					return db, nil
				}
				err = fmt.Errorf("ping failed: %w", pingErr)
			} else {
				err = fmt.Errorf("getting raw DB handle failed: %w", pingErr)
			}
		}
		log.Printf("⚠️  Attempt %d/%d to connect database failed: %v", attempt, maxAttempts, err)
		time.Sleep(delay)
	}
	return nil, err
}


// @title Documents Service API
// @version 1.0
// @description This is a API for managing documents.

// @host localhost:8080
// @BasePath /

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Connect to database using config
	db, err := connectDatabase(cfg.GetDSN(), 10, 2*time.Second) // 10 attempts, 2s apart
	if err != nil {
		log.Fatalf("❌ Could not connect to database after retries: %v", err)
	}


	//database migration
	if err := db.AutoMigrate(&models.DocumentFromOrm{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database connected successfully")

	repo := Repositories.NewGormRepository[models.Document](db)
	docService := Services.NewDocumentService(repo)
	docHandler := Handlers.NewDocumentHandler(docService)

	// Update SwaggerInfo
	docs.SwaggerInfo.Host = cfg.SwaggerHost

	r := routes.SetupRouter(docHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Use config port
	log.Printf("🚀 Starting server on :%s", cfg.APIPort)
	if err := r.Run(":" + cfg.APIPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

