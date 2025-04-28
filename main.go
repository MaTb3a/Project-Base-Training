package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connectDatabase attempts to open a GORM connection and ping the DB.
// It retries up to maxAttempts times, waiting delay between tries.
func connectDatabase(dsn string, maxAttempts int, delay time.Duration) (*gorm.DB, error) {
    var db *gorm.DB
    var err error

    for attempt := 1; attempt <= maxAttempts; attempt++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

    return nil, err
}

func main() {
    // Build the Postgres DSN from environment
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        getenv("DB_HOST", "db"),
        getenv("DB_USER", "postgres"),
        getenv("DB_PASSWORD", "postgres"),
        getenv("DB_NAME", "document_db"),
        getenv("DB_PORT", "5432"),
        getenv("DB_SSL_MODE", "disable"),
    )

    // Try connecting up to 10 times, waiting 2s between attempts
    _, err := connectDatabase(dsn, 10, 2*time.Second)
    if err != nil {
        log.Fatalf("❌ Could not connect to database after retries: %v", err)
    }
    // (You can now use `db` for migrations, queries, etc.)

    // Minimal Gin setup with a /ping endpoint
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

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
