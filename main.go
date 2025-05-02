package main

import (
	"fmt"
	"log"
	"os"
	"time"

	Repositories "github.com/MaTb3aa/Project-Base-Training/repository"
	"github.com/MaTb3aa/Project-Base-Training/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	Handlers "github.com/MaTb3aa/Project-Base-Training/handdlers"
	"github.com/MaTb3aa/Project-Base-Training/models"
	Services "github.com/MaTb3aa/Project-Base-Training/services"
)

// connectDatabase attempts to open a GORM connection and ping the DB.
// It retries up to maxAttempts times, waiting delay between tries.
func connectDatabase(dsn string, maxAttempts int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open("host=localhost user=postgres password=123 dbname=document_db port=5432 sslmode=disable"), &gorm.Config{})

		//database migration
		if err := db.AutoMigrate(&models.DocumentFromOrm{}); err != nil {
			log.Fatal("Migration failed:", err)
		}
		if err == nil {
			// Verify lower-level connectivity
			sqlDB, pingErr := db.DB()
			if pingErr == nil {
				if pingErr = sqlDB.Ping(); pingErr == nil {
					log.Printf("✅ Database connected on attempt %d", attempt)
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
	db.AutoMigrate(&models.DocumentFromOrm{})

	return nil, err
}

func main() {
	//Build the Postgres DSN from environment
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getenv("DB_HOST", "localhost"),
		getenv("DB_USER", "postgres"),
		getenv("DB_PASSWORD", "postgres"),
		getenv("DB_NAME", "document_db"),
		getenv("DB_PORT", "5432"),
		getenv("DB_SSL_MODE", "disable"),
	)

	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}
	//database migration
	if err := db.AutoMigrate(&models.DocumentFromOrm{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database connected successfully")

	repo := Repositories.NewGormRepository[models.Document](db)
	docService := Services.NewDocumentService(repo)
	docHandler := Handlers.NewDocumentHandler(docService)

	r := routes.SetupRouter(docHandler)
	port := getenv("PORT", "8080")
	log.Printf("🚀 Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// getenv returns the value of the environment variable named by the key,
// or fallback if the variable is empty or not present.
func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
