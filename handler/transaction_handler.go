package handler

import (
	"encoding/json"
	"kasir-api/entity"
	"kasir-api/service"
	"net/http"
)

// TransactionHandler - struct untuk transaction handler
type TransactionHandler struct {
	service service.TransactionServiceInterface
}

// NewTransactionHandler - constructor untuk TransactionHandler
func NewTransactionHandler(service service.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// HandleCheckout - route handler untuk /api/checkout
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Checkout - handler untuk POST /api/checkout
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req entity.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate items
	if len(req.Items) == 0 {
		http.Error(w, "Items cannot be empty", http.StatusBadRequest)
		return
	}

	for _, item := range req.Items {
		if item.ProductID <= 0 {
			http.Error(w, "Invalid product_id", http.StatusBadRequest)
			return
		}
		if item.Quantity <= 0 {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// HandleSalesReportToday - route handler untuk /api/report/hari-ini
func (h *TransactionHandler) HandleSalesReportToday(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetSalesReportToday(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetSalesReportToday - handler untuk GET /api/report/hari-ini
func (h *TransactionHandler) GetSalesReportToday(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetSalesReportToday()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// HandleSalesReport - route handler untuk /api/report dengan query params
func (h *TransactionHandler) HandleSalesReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetSalesReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetSalesReport - handler untuk GET /api/report?start_date=2026-01-01&end_date=2026-02-01
func (h *TransactionHandler) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// If no date range provided, return today's report with extended format
	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date query parameters are required", http.StatusBadRequest)
		return
	}

	report, err := h.service.GetSalesReportByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
