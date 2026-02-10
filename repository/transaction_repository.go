package repository

import (
	"database/sql"
	"fmt"
	"kasir-api/entity"
	"time"
)

// TransactionRepositoryInterface - interface untuk transaction repository
type TransactionRepositoryInterface interface {
	CreateTransaction(items []entity.CheckoutItem) (*entity.Transaction, error)
	GetSalesReportToday() (*entity.SalesReport, error)
	GetSalesReportByDateRange(startDate, endDate string) (*entity.SalesReportWithRange, error)
}

// TransactionRepository - struct untuk transaction repository
type TransactionRepository struct {
	db *sql.DB
}

// NewTransactionRepository - constructor untuk TransactionRepository
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// CreateTransaction - membuat transaksi baru dengan multiple items
func (r *TransactionRepository) CreateTransaction(items []entity.CheckoutItem) (*entity.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]entity.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT nama, harga, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Check if stock is sufficient
		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", productName, stock, item.Quantity)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// Update stock
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, entity.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Insert transaction
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Batch insert transaction details - FIXED: using copy or prepared statement for efficiency
	stmt, err := tx.Prepare("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = stmt.Exec(transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &entity.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

// GetSalesReportToday - mendapatkan laporan penjualan hari ini
func (r *TransactionRepository) GetSalesReportToday() (*entity.SalesReport, error) {
	today := time.Now().Format("2006-01-02")

	// Get total revenue and transaction count
	var totalRevenue, totalTransaksi int
	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = $1
	`, today).Scan(&totalRevenue, &totalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	var produkTerlaris entity.ProdukLaris
	row := r.db.QueryRow(`
		SELECT p.nama, SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = $1
		GROUP BY p.id, p.nama
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, today)

	var nama string
	var qtyTerjual int
	scanErr := row.Scan(&nama, &qtyTerjual)
	if scanErr == nil {
		produkTerlaris = entity.ProdukLaris{
			Nama:       nama,
			QtyTerjual: qtyTerjual,
		}
	} else {
		// No sales today
		produkTerlaris = entity.ProdukLaris{
			Nama:       "-",
			QtyTerjual: 0,
		}
	}

	return &entity.SalesReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}, nil
}

// GetSalesReportByDateRange - mendapatkan laporan penjualan berdasarkan range tanggal
func (r *TransactionRepository) GetSalesReportByDateRange(startDate, endDate string) (*entity.SalesReportWithRange, error) {
	// Get total revenue and transaction count
	var totalRevenue, totalTransaksi int
	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`, startDate, endDate).Scan(&totalRevenue, &totalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	var produkTerlaris entity.ProdukLaris
	row := r.db.QueryRow(`
		SELECT p.nama, SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY p.id, p.nama
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, startDate, endDate)

	var nama string
	var qtyTerjual int
	scanErr := row.Scan(&nama, &qtyTerjual)
	if scanErr == nil {
		produkTerlaris = entity.ProdukLaris{
			Nama:       nama,
			QtyTerjual: qtyTerjual,
		}
	} else {
		// No sales in date range
		produkTerlaris = entity.ProdukLaris{
			Nama:       "-",
			QtyTerjual: 0,
		}
	}

	return &entity.SalesReportWithRange{
		StartDate:      startDate,
		EndDate:        endDate,
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}, nil
}
