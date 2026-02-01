package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"kasir-api/config/migration"
	"kasir-api/config/seeder"

	_ "github.com/lib/pq"
)

// DB is the global database connection
var DB *sql.DB

// ConnectDB initializes database connection using URI
func ConnectDB() *sql.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Extract database name from URL
	dbName := extractDBName(databaseURL)
	
	// Connect to default postgres database first to create target database
	postgresURL := replaceDBName(databaseURL, "postgres")
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Fatal("Failed to connect to postgres database:", err)
	}
	
	// Create database if not exists
	_, _ = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	db.Close()

	// Connect to target database
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL database")
	DB = db
	return db
}

// SetupDatabase runs migrations and seeders
func SetupDatabase() {
	if DB == nil {
		log.Fatal("Database not connected. Call ConnectDB() first.")
	}

	// Run migrations
	migration.RunMigration(DB)

	// Run seeders
	seeder.RunSeeder(DB)
}

// extractDBName extracts database name from connection URL
func extractDBName(url string) string {
	// Format: postgres://user:pass@host:port/dbname?options
	parts := strings.Split(url, "/")
	if len(parts) >= 4 {
		dbPart := parts[3]
		// Remove query parameters
		if idx := strings.Index(dbPart, "?"); idx != -1 {
			return dbPart[:idx]
		}
		return dbPart
	}
	return "kasir_api"
}

// replaceDBName replaces database name in connection URL
func replaceDBName(url, newDB string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 4 {
		// Replace database name
		dbPart := parts[3]
		if idx := strings.Index(dbPart, "?"); idx != -1 {
			parts[3] = newDB + dbPart[idx:]
		} else {
			parts[3] = newDB
		}
		return strings.Join(parts, "/")
	}
	return url
}
