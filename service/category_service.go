package service

import (
	"kasir-api/entity"
	"kasir-api/repository"
)

// CategoryServiceInterface - interface untuk category service
type CategoryServiceInterface interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id int) (entity.Category, error)
	CreateCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(id int, category entity.Category) (entity.Category, error)
	DeleteCategory(id int) error
}

// CategoryService - struct untuk category service
type CategoryService struct {
	repo repository.CategoryRepositoryInterface
}

// NewCategoryService - constructor untuk CategoryService
func NewCategoryService(repo repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAllCategories - ambil semua kategori
func (s *CategoryService) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAll()
}

// GetCategoryByID - ambil kategori berdasarkan ID
func (s *CategoryService) GetCategoryByID(id int) (entity.Category, error) {
	return s.repo.GetByID(id)
}

// CreateCategory - tambah kategori baru
func (s *CategoryService) CreateCategory(category entity.Category) (entity.Category, error) {
	return s.repo.Create(category)
}

// UpdateCategory - update kategori
func (s *CategoryService) UpdateCategory(id int, category entity.Category) (entity.Category, error) {
	return s.repo.Update(id, category)
}

// DeleteCategory - hapus kategori
func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}
