package repository

import (
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
	categories []entity.Category
}

// NewCategoryRepository - constructor untuk CategoryRepository
func NewCategoryRepository() *CategoryRepository {
	// Data awal
	return &CategoryRepository{
		categories: []entity.Category{
			{ID: 1, Name: "Minuman", Description: "Segala jenis minuman"},
			{ID: 2, Name: "Makanan", Description: "Segala jenis makanan"},
		},
	}
}

// GetAll - ambil semua kategori
func (r *CategoryRepository) GetAll() ([]entity.Category, error) {
	return r.categories, nil
}

// GetByID - ambil kategori berdasarkan ID
func (r *CategoryRepository) GetByID(id int) (entity.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return c, nil
		}
	}
	return entity.Category{}, errors.New("category not found")
}

// Create - tambah kategori baru
func (r *CategoryRepository) Create(category entity.Category) (entity.Category, error) {
	category.ID = len(r.categories) + 1
	r.categories = append(r.categories, category)
	return category, nil
}

// Update - update kategori
func (r *CategoryRepository) Update(id int, category entity.Category) (entity.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			category.ID = id
			r.categories[i] = category
			return category, nil
		}
	}
	return entity.Category{}, errors.New("category not found")
}

// Delete - hapus kategori
func (r *CategoryRepository) Delete(id int) error {
	for i, c := range r.categories {
		if c.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}
