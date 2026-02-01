package seeder

import (
	"database/sql"
	"fmt"
)

// CategorySeed represents a category seed data
type CategorySeed struct {
	Name        string
	Description string
}

// DefaultCategories contains default category data
var DefaultCategories = []CategorySeed{
	{
		Name:        "Minuman",
		Description: "Segala jenis minuman",
	},
	{
		Name:        "Makanan",
		Description: "Segala jenis makanan",
	},
	{
		Name:        "Snack",
		Description: "Makanan ringan dan cemilan",
	},
	{
		Name:        "Elektronik",
		Description: "Barang elektronik dan gadget",
	},
}

// SeedCategories seeds categories data
func SeedCategories(db *sql.DB) error {
	fmt.Println("🌱 Seeding categories...")

	// Check if already seeded
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check categories count: %w", err)
	}

	if count > 0 {
		fmt.Printf("  ⏭️  Categories already seeded (%d records exist)\n", count)
		return nil
	}

	// Insert categories
	for _, cat := range DefaultCategories {
		_, err := db.Exec(
			"INSERT INTO categories (name, description) VALUES ($1, $2)",
			cat.Name, cat.Description,
		)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", cat.Name, err)
		}
		fmt.Printf("  ✓ Category: %s\n", cat.Name)
	}

	fmt.Printf("  ✅ Seeded %d categories\n", len(DefaultCategories))
	return nil
}

// GetCategoryIDByName gets category ID by name
func GetCategoryIDByName(db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM categories WHERE name = $1", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
