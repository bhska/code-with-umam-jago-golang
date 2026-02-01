package repository

import (
	"errors"
	"kasir-api/entity"
)

// ProductRepositoryInterface - interface untuk product repository
type ProductRepositoryInterface interface {
	GetAll() ([]entity.Product, error)
	GetByID(id int) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	Update(id int, product entity.Product) (entity.Product, error)
	Delete(id int) error
}

// ProductRepository - struct untuk product repository
type ProductRepository struct {
	products []entity.Product
}

// NewProductRepository - constructor untuk ProductRepository
func NewProductRepository() *ProductRepository {
	// Data awal
	return &ProductRepository{
		products: []entity.Product{
			{ID: 1, Nama: "Produk A", Harga: 10000, CategoryID: 1},
			{ID: 2, Nama: "Produk B", Harga: 20000, CategoryID: 2},
			{ID: 3, Nama: "Produk C", Harga: 30000, CategoryID: 1},
		},
	}
}

// GetAll - ambil semua produk
func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	return r.products, nil
}

// GetByID - ambil produk berdasarkan ID
func (r *ProductRepository) GetByID(id int) (entity.Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return entity.Product{}, errors.New("product not found")
}

// Create - tambah produk baru
func (r *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return product, nil
}

// Update - update produk
func (r *ProductRepository) Update(id int, product entity.Product) (entity.Product, error) {
	for i := range r.products {
		if r.products[i].ID == id {
			product.ID = id
			r.products[i] = product
			return product, nil
		}
	}
	return entity.Product{}, errors.New("product not found")
}

// Delete - hapus produk
func (r *ProductRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
