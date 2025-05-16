package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    // API Settings
    APIPort  string
    GinMode  string

    // Database Settings
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string

    // Swagger Settings
    SwaggerHost string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: .env file not found, using default values")
    }

    return &Config{
        // API Settings
        APIPort:  getEnv("API_PORT", "8888"),
        GinMode:  getEnv("GIN_MODE", "debug"),

        // Database Settings
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "document_db"),
        DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

        // Swagger Settings
        SwaggerHost: getEnv("SWAGGER_HOST", "localhost:8888"),
    }
}

func (c *Config) GetDSN() string {
    return fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode,
    )
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}