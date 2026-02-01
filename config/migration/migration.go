package migration

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed *.sql
var migrationFiles embed.FS

// Migration represents a single migration
type Migration struct {
	Version string
	Name    string
	SQL     string
}

// Migrate runs all pending migrations
func Migrate(db *sql.DB) error {
	fmt.Println("🔄 Running migrations...")

	// Create migrations tracking table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get list of migration files
	files, err := migrationFiles.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	// Sort files by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Execute each migration
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		version := strings.TrimSuffix(file.Name(), ".sql")

		// Check if migration already applied
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)", version).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check migration status: %w", err)
		}

		if exists {
			fmt.Printf("  ✓ %s (already applied)\n", version)
			continue
		}

		// Read migration file
		content, err := migrationFiles.ReadFile(file.Name())
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
		}

		// Execute migration in transaction
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		_, err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute migration %s: %w", version, err)
		}

		_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", version, err)
		}

		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", version, err)
		}

		fmt.Printf("  ✓ %s (applied)\n", version)
	}

	fmt.Println("✅ All migrations completed")
	return nil
}

// Rollback rolls back the last migration
func Rollback(db *sql.DB) error {
	fmt.Println("🔄 Rolling back last migration...")

	// Get last applied migration
	var version string
	err := db.QueryRow("SELECT version FROM schema_migrations ORDER BY applied_at DESC LIMIT 1").Scan(&version)
	if err == sql.ErrNoRows {
		fmt.Println("  No migrations to rollback")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to get last migration: %w", err)
	}

	// Read migration file for rollback SQL (if exists)
	rollbackFile := version + "_rollback.sql"
	content, err := migrationFiles.ReadFile(rollbackFile)
	if err != nil {
		// If no rollback file, just delete the migration record
		_, err = db.Exec("DELETE FROM schema_migrations WHERE version = $1", version)
		if err != nil {
			return fmt.Errorf("failed to remove migration record: %w", err)
		}
		fmt.Printf("  ✓ %s (removed record only, no rollback file)\n", version)
		return nil
	}

	// Execute rollback
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec(string(content))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute rollback: %w", err)
	}

	_, err = tx.Exec("DELETE FROM schema_migrations WHERE version = $1", version)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to remove migration record: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit rollback: %w", err)
	}

	fmt.Printf("  ✓ %s (rolled back)\n", version)
	return nil
}

// Status shows current migration status
func Status(db *sql.DB) error {
	fmt.Println("📊 Migration Status:")

	rows, err := db.Query("SELECT version, applied_at FROM schema_migrations ORDER BY applied_at")
	if err != nil {
		return fmt.Errorf("failed to get migration status: %w", err)
	}
	defer rows.Close()

	fmt.Println("  Applied migrations:")
	for rows.Next() {
		var version string
		var appliedAt string
		err := rows.Scan(&version, &appliedAt)
		if err != nil {
			return err
		}
		fmt.Printf("    ✓ %s (applied at %s)\n", version, appliedAt)
	}

	return nil
}

// GetFilename extracts filename without extension for version
func GetFilename(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// RunMigration runs migrations
func RunMigration(db *sql.DB) {
	if err := Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}
}
