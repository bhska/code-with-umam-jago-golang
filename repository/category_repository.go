package repository

import (
	"database/sql"
	"errors"
	"kasir-api/entity"
)

// CategoryRepositoryInterface - interface untuk category repository
type CategoryRepositoryInterface interface {
	GetAll() ([]entity.Category, error)
	GetByID(id int) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
	Update(id int, category entity.Category) (entity.Category, error)
	Delete(id int) error
}

// CategoryRepository - struct untuk category repository
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository - constructor untuk CategoryRepository
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll - ambil semua kategori
func (r *CategoryRepository) GetAll() ([]entity.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var c entity.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

// GetByID - ambil kategori berdasarkan ID
func (r *CategoryRepository) GetByID(id int) (entity.Category, error) {
	var c entity.Category
	err := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&c.ID, &c.Name, &c.Description)
	
	if err == sql.ErrNoRows {
		return entity.Category{}, errors.New("category not found")
	}
	if err != nil {
		return entity.Category{}, err
	}
	
	return c, nil
}

// Create - tambah kategori baru
func (r *CategoryRepository) Create(category entity.Category) (entity.Category, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		category.Name, category.Description,
	).Scan(&id)
	
	if err != nil {
		return entity.Category{}, err
	}
	
	category.ID = id
	return category, nil
}

// Update - update kategori
func (r *CategoryRepository) Update(id int, category entity.Category) (entity.Category, error) {
	result, err := r.db.Exec(
		"UPDATE categories SET name = $1, description = $2 WHERE id = $3",
		category.Name, category.Description, id,
	)
	if err != nil {
		return entity.Category{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.Category{}, err
	}
	if rowsAffected == 0 {
		return entity.Category{}, errors.New("category not found")
	}

	category.ID = id
	return category, nil
}

// Delete - hapus kategori
func (r *CategoryRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}
