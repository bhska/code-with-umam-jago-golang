package seeder

import (
	"database/sql"
	"fmt"
)

// ProductSeed represents a product seed data
type ProductSeed struct {
	Nama       string
	Harga      int
	CategoryID int
}

// DefaultProducts contains default product data
var DefaultProducts = []ProductSeed{
	{
		Nama:       "Es Teh Manis",
		Harga:      5000,
		CategoryID: 1, // Minuman
	},
	{
		Nama:       "Kopi Hitam",
		Harga:      8000,
		CategoryID: 1, // Minuman
	},
	{
		Nama:       "Nasi Goreng",
		Harga:      15000,
		CategoryID: 2, // Makanan
	},
	{
		Nama:       "Mie Ayam",
		Harga:      12000,
		CategoryID: 2, // Makanan
	},
	{
		Nama:       "Keripik Kentang",
		Harga:      8000,
		CategoryID: 3, // Snack
	},
	{
		Nama:       "Chocolatos",
		Harga:      2000,
		CategoryID: 3, // Snack
	},
}

// SeedProducts seeds products data
func SeedProducts(db *sql.DB) error {
	fmt.Println("🌱 Seeding products...")

	// Check if already seeded
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check products count: %w", err)
	}

	if count > 0 {
		fmt.Printf("  ⏭️  Products already seeded (%d records exist)\n", count)
		return nil
	}

	// Insert products
	for _, prod := range DefaultProducts {
		_, err := db.Exec(
			"INSERT INTO products (nama, harga, category_id) VALUES ($1, $2, $3)",
			prod.Nama, prod.Harga, prod.CategoryID,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product %s: %w", prod.Nama, err)
		}
		fmt.Printf("  ✓ Product: %s (Rp %d)\n", prod.Nama, prod.Harga)
	}

	fmt.Printf("  ✅ Seeded %d products\n", len(DefaultProducts))
	return nil
}

// SeedProductsWithCategoryNames seeds products using category names
func SeedProductsWithCategoryNames(db *sql.DB) error {
	fmt.Println("🌱 Seeding products with category lookup...")

	// Check if already seeded
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check products count: %w", err)
	}

	if count > 0 {
		fmt.Printf("  ⏭️  Products already seeded (%d records exist)\n", count)
		return nil
	}

	// Product definitions with category names
	products := []struct {
		Nama         string
		Harga        int
		CategoryName string
	}{
		{"Es Teh Manis", 5000, "Minuman"},
		{"Kopi Hitam", 8000, "Minuman"},
		{"Nasi Goreng", 15000, "Makanan"},
		{"Mie Ayam", 12000, "Makanan"},
		{"Keripik Kentang", 8000, "Snack"},
		{"Chocolatos", 2000, "Snack"},
	}

	// Insert products with category lookup
	for _, prod := range products {
		// Get category ID
		var categoryID int
		err := db.QueryRow("SELECT id FROM categories WHERE name = $1", prod.CategoryName).Scan(&categoryID)
		if err != nil {
			fmt.Printf("  ⚠️  Skipping %s: category '%s' not found\n", prod.Nama, prod.CategoryName)
			continue
		}

		_, err = db.Exec(
			"INSERT INTO products (nama, harga, category_id) VALUES ($1, $2, $3)",
			prod.Nama, prod.Harga, categoryID,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product %s: %w", prod.Nama, err)
		}
		fmt.Printf("  ✓ Product: %s (Rp %d) - %s\n", prod.Nama, prod.Harga, prod.CategoryName)
	}

	fmt.Println("  ✅ Products seeded")
	return nil
}
