package main

import (
	"encoding/json"
	_ "kasir-api/docs"
	"kasir-api/handler"
	"kasir-api/repository"
	"kasir-api/service"
	"net/http"
	"os"

	"kasir-api/config"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// APIInfo represents the API information for root endpoint
type APIInfo struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	Description string       `json:"description"`
	Database    string       `json:"database"`
	Endpoints   Endpoints    `json:"endpoints"`
	Architecture Architecture `json:"architecture"`
}

// Endpoints represents all available endpoints
type Endpoints struct {
	Root        string `json:"root"`
	Health      string `json:"health"`
	Swagger     string `json:"swagger"`
	Categories  string `json:"categories"`
	Products    string `json:"products"`
	Checkout    string `json:"checkout"`
	ReportToday string `json:"report_today"`
	ReportRange string `json:"report_range"`
}

// Architecture represents the layered architecture
type Architecture struct {
	Layers []Layer `json:"layers"`
}

// Layer represents a single layer
type Layer struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func getAPIInfo() APIInfo {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	
	baseURL := "http://localhost:" + port
	
	return APIInfo{
		Name:        "Kasir API",
		Version:     "1.0.0",
		Description: "API Kasir dengan Layered Architecture - Week 2 Challenge",
		Database:    "PostgreSQL (Neon)",
		Endpoints: Endpoints{
			Root:        baseURL + "/",
			Health:      baseURL + "/health",
			Swagger:     baseURL + "/swagger/",
			Categories:  baseURL + "/api/categories",
			Products:    baseURL + "/api/produk",
			Checkout:    baseURL + "/api/checkout",
			ReportToday: baseURL + "/api/report/hari-ini",
			ReportRange: baseURL + "/api/report?start_date=2026-01-01&end_date=2026-02-01",
		},
		Architecture: Architecture{
			Layers: []Layer{
				{Name: "Handler", Path: "handler/", Description: "HTTP Layer - Request/Response handling"},
				{Name: "Service", Path: "service/", Description: "Business Logic Layer"},
				{Name: "Repository", Path: "repository/", Description: "Data Access Layer (PostgreSQL/Neon)"},
				{Name: "Entity", Path: "entity/", Description: "Models/Entities"},
			},
		},
	}
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		println("⚠️  Warning: .env file not found, using environment variables")
	}

	// Connect to PostgreSQL Database (Neon)
	db := config.ConnectDB()
	defer db.Close()

	// Run migrations and seeders
	config.SetupDatabase()

	// ===== LAYERED ARCHITECTURE SETUP =====
	
	// Repository Layer (Data Access with PostgreSQL/Neon)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	
	// Service Layer (Business Logic)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo, categoryRepo) // Inject categoryRepo untuk JOIN
	transactionService := service.NewTransactionService(transactionRepo)
	
	// Handler Layer (HTTP Handler/Controller)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	
	// ===== ROUTES =====
	
	// Root endpoint - Simple JSON
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getAPIInfo())
	})
	
	// Swagger UI
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	
	// Category Routes (Layered Architecture)
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetCategoryByID(w, r)
		case "PUT":
			categoryHandler.UpdateCategory(w, r)
		case "DELETE":
			categoryHandler.DeleteCategory(w, r)
		}
	})
	
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			categoryHandler.GetAllCategories(w, r)
		case "POST":
			categoryHandler.CreateCategory(w, r)
		}
	})
	
	// Product Routes (Layered Architecture dengan CHALLENGE: JOIN)
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// CHALLENGE: Get Detail Product dengan Category Name (JOIN)
			productHandler.GetProductByID(w, r)
		case "PUT":
			productHandler.UpdateProduct(w, r)
		case "DELETE":
			productHandler.DeleteProduct(w, r)
		}
	})
	
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetAllProducts(w, r)
		case "POST":
			productHandler.CreateProduct(w, r)
		}
	})
	
	// Transaction Routes
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST
	http.HandleFunc("/api/report/hari-ini", transactionHandler.HandleSalesReportToday) // GET
	http.HandleFunc("/api/report", transactionHandler.HandleSalesReport) // GET with query params

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"OK","message":"API Running with PostgreSQL (Neon)"}`))
	})
	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	
	// Server info
	println("╔════════════════════════════════════════════════════════════╗")
	println("║                    🚀 Kasir API Server                     ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  📍 Endpoint Utama:  http://localhost:" + port + "/                ║")
	println("║  📖 Swagger UI:      http://localhost:" + port + "/swagger/        ║")
	println("║  💓 Health Check:    http://localhost:" + port + "/health          ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  🐘 Database: PostgreSQL (Neon Serverless)                 ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  Layered Architecture:                                     ║")
	println("║    • Entity     → entity/                                  ║")
	println("║    • Repository → repository/ (PostgreSQL)                 ║")
	println("║    • Service    → service/                                 ║")
	println("║    • Handler    → handler/                                 ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  ✨ Session 3: Transaction API & Sales Report              ║")
	println("║     • POST /api/checkout - Create transaction              ║")
	println("║     • GET  /api/report/hari-ini - Today's sales report     ║")
	println("║     • GET  /api/report?start_date=&end_date= - Date range  ║")
	println("║                                                            ║")
	println("║  ⭐ Challenge: JOIN Product dengan Category!                ║")
	println("╚════════════════════════════════════════════════════════════╝")
	
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		println("❌ Gagal running server:", err.Error())
	}
}
