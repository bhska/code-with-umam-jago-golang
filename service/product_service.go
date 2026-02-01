package service

import (
	"kasir-api/entity"
	"kasir-api/repository"
)

// ProductServiceInterface - interface untuk product service
type ProductServiceInterface interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id int) (entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(id int, product entity.Product) (entity.Product, error)
	DeleteProduct(id int) error
}

// ProductService - struct untuk product service
type ProductService struct {
	productRepo  repository.ProductRepositoryInterface
	categoryRepo repository.CategoryRepositoryInterface
}

// NewProductService - constructor untuk ProductService
func NewProductService(productRepo repository.ProductRepositoryInterface, categoryRepo repository.CategoryRepositoryInterface) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// GetAllProducts - ambil semua produk
func (s *ProductService) GetAllProducts() ([]entity.Product, error) {
	return s.productRepo.GetAll()
}

// GetProductByID - ambil produk berdasarkan ID dengan join category
func (s *ProductService) GetProductByID(id int) (entity.Product, error) {
	// Ambil produk dari repository
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return entity.Product{}, err
	}

	// Join: Ambil category berdasarkan CategoryID
	category, err := s.categoryRepo.GetByID(product.CategoryID)
	if err == nil {
		// Jika category ditemukan, tambahkan ke product
		product.Category = &category
	}

	return product, nil
}

// CreateProduct - tambah produk baru
func (s *ProductService) CreateProduct(product entity.Product) (entity.Product, error) {
	return s.productRepo.Create(product)
}

// UpdateProduct - update produk
func (s *ProductService) UpdateProduct(id int, product entity.Product) (entity.Product, error) {
	return s.productRepo.Update(id, product)
}

// DeleteProduct - hapus produk
func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepo.Delete(id)
}
