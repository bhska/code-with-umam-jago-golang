package service

import (
	"kasir-api/entity"
	"kasir-api/repository"
)

// TransactionServiceInterface - interface untuk transaction service
type TransactionServiceInterface interface {
	Checkout(items []entity.CheckoutItem) (*entity.Transaction, error)
	GetSalesReportToday() (*entity.SalesReport, error)
	GetSalesReportByDateRange(startDate, endDate string) (*entity.SalesReportWithRange, error)
}

// TransactionService - struct untuk transaction service
type TransactionService struct {
	transactionRepo repository.TransactionRepositoryInterface
}

// NewTransactionService - constructor untuk TransactionService
func NewTransactionService(transactionRepo repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

// Checkout - memproses checkout transaksi
func (s *TransactionService) Checkout(items []entity.CheckoutItem) (*entity.Transaction, error) {
	return s.transactionRepo.CreateTransaction(items)
}

// GetSalesReportToday - mendapatkan laporan penjualan hari ini
func (s *TransactionService) GetSalesReportToday() (*entity.SalesReport, error) {
	return s.transactionRepo.GetSalesReportToday()
}

// GetSalesReportByDateRange - mendapatkan laporan penjualan berdasarkan range tanggal
func (s *TransactionService) GetSalesReportByDateRange(startDate, endDate string) (*entity.SalesReportWithRange, error) {
	return s.transactionRepo.GetSalesReportByDateRange(startDate, endDate)
}
