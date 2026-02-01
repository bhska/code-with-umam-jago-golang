package seeder

import (
	"database/sql"
	"fmt"
	"log"
)

// Seeder handles database seeding
type Seeder struct {
	db *sql.DB
}

// NewSeeder creates a new seeder instance
func NewSeeder(db *sql.DB) *Seeder {
	return &Seeder{db: db}
}

// Run runs all seeders
func (s *Seeder) Run() error {
	fmt.Println("\n🌱 Running database seeders...")

	// Run seeders in order
	if err := SeedCategories(s.db); err != nil {
		return fmt.Errorf("category seeder failed: %w", err)
	}

	if err := SeedProductsWithCategoryNames(s.db); err != nil {
		return fmt.Errorf("product seeder failed: %w", err)
	}

	fmt.Println("✅ All seeders completed successfully")
	return nil
}

// RunSeeder runs the seeder
func RunSeeder(db *sql.DB) {
	seeder := NewSeeder(db)
	if err := seeder.Run(); err != nil {
		log.Fatal("Seeder failed:", err)
	}
}

// Clear clears all data (useful for testing)
func Clear(db *sql.DB) error {
	fmt.Println("🗑️  Clearing all data...")

	// Delete in correct order (products first due to FK)
	_, err := db.Exec("DELETE FROM products")
	if err != nil {
		return fmt.Errorf("failed to clear products: %w", err)
	}
	fmt.Println("  ✓ Cleared products")

	_, err = db.Exec("DELETE FROM categories")
	if err != nil {
		return fmt.Errorf("failed to clear categories: %w", err)
	}
	fmt.Println("  ✓ Cleared categories")

	// Reset sequences
	_, err = db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
	if err != nil {
		return fmt.Errorf("failed to reset products sequence: %w", err)
	}

	_, err = db.Exec("ALTER SEQUENCE categories_id_seq RESTART WITH 1")
	if err != nil {
		return fmt.Errorf("failed to reset categories sequence: %w", err)
	}

	fmt.Println("✅ All data cleared")
	return nil
}

// Refresh clears and re-seeds data
func Refresh(db *sql.DB) error {
	if err := Clear(db); err != nil {
		return err
	}
	return NewSeeder(db).Run()
}
