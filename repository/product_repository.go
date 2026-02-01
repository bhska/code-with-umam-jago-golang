package repository

import (
	"database/sql"
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
	db *sql.DB
}

// NewProductRepository - constructor untuk ProductRepository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll - ambil semua produk
func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	rows, err := r.db.Query("SELECT id, nama, harga, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// GetByID - ambil produk berdasarkan ID
func (r *ProductRepository) GetByID(id int) (entity.Product, error) {
	var p entity.Product
	err := r.db.QueryRow(
		"SELECT id, nama, harga, category_id FROM products WHERE id = $1", id,
	).Scan(&p.ID, &p.Nama, &p.Harga, &p.CategoryID)
	
	if err == sql.ErrNoRows {
		return entity.Product{}, errors.New("product not found")
	}
	if err != nil {
		return entity.Product{}, err
	}
	
	return p, nil
}

// Create - tambah produk baru
func (r *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO products (nama, harga, category_id) VALUES ($1, $2, $3) RETURNING id",
		product.Nama, product.Harga, product.CategoryID,
	).Scan(&id)
	
	if err != nil {
		return entity.Product{}, err
	}
	
	product.ID = id
	return product, nil
}

// Update - update produk
func (r *ProductRepository) Update(id int, product entity.Product) (entity.Product, error) {
	result, err := r.db.Exec(
		"UPDATE products SET nama = $1, harga = $2, category_id = $3 WHERE id = $4",
		product.Nama, product.Harga, product.CategoryID, id,
	)
	if err != nil {
		return entity.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.Product{}, err
	}
	if rowsAffected == 0 {
		return entity.Product{}, errors.New("product not found")
	}

	product.ID = id
	return product, nil
}

// Delete - hapus produk
func (r *ProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
