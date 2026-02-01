package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL database (Neon)")
	DB = db
	return db
}

// Migrate creates tables if not exist
func Migrate() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	// Create categories table
	categoriesSQL := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT
	);`

	_, err := DB.Exec(categoriesSQL)
	if err != nil {
		log.Fatal("Failed to create categories table:", err)
	}

	// Create products table
	productsSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		harga INTEGER NOT NULL,
		category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
	);`

	_, err = DB.Exec(productsSQL)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	fmt.Println("✅ Database migrated successfully")

	// Seed default data
	seedData()
}

// seedData inserts default data
func seedData() {
	// Check if categories already exist
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		log.Println("Failed to check categories count:", err)
		return
	}

	if count > 0 {
		fmt.Println("📦 Data already seeded")
		return
	}

	// Insert default categories
	_, err = DB.Exec(`
		INSERT INTO categories (name, description) VALUES
		('Minuman', 'Segala jenis minuman'),
		('Makanan', 'Segala jenis makanan');
	`)
	if err != nil {
		log.Println("Failed to seed categories:", err)
		return
	}

	// Insert default products
	_, err = DB.Exec(`
		INSERT INTO products (nama, harga, category_id) VALUES
		('Produk A', 10000, 1),
		('Produk B', 20000, 2),
		('Produk C', 30000, 1);
	`)
	if err != nil {
		log.Println("Failed to seed products:", err)
		return
	}

	fmt.Println("🌱 Default data seeded")
}
