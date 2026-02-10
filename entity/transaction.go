package entity

import "time"

// Transaction represents a sale transaction
type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

// TransactionDetail represents items in a transaction
type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

// CheckoutItem represents an item in checkout request
type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// CheckoutRequest represents the checkout request body
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

// SalesReport represents daily sales report
type SalesReport struct {
	TotalRevenue   int           `json:"total_revenue"`
	TotalTransaksi int           `json:"total_transaksi"`
	ProdukTerlaris ProdukLaris   `json:"produk_terlaris"`
}

// ProdukLaris represents best selling product info
type ProdukLaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

// SalesReportWithRange represents sales report with date range
type SalesReportWithRange struct {
	StartDate      string        `json:"start_date"`
	EndDate        string        `json:"end_date"`
	TotalRevenue   int           `json:"total_revenue"`
	TotalTransaksi int           `json:"total_transaksi"`
	ProdukTerlaris ProdukLaris   `json:"produk_terlaris"`
}
