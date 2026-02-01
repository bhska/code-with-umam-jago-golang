package main

import (
	_ "kasir-api/docs"
	"kasir-api/handler"
	"kasir-api/repository"
	"kasir-api/service"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// HTML untuk halaman root /
const indexHTML = `<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kasir API - Week 2</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            padding: 40px;
            max-width: 800px;
            width: 100%;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
        }
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 10px;
            font-size: 2.5em;
        }
        .subtitle {
            text-align: center;
            color: #666;
            margin-bottom: 30px;
        }
        .badge {
            display: inline-block;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 5px 15px;
            border-radius: 20px;
            font-size: 0.85em;
            margin: 5px;
        }
        .section {
            margin: 25px 0;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 10px;
        }
        .section h2 {
            color: #667eea;
            margin-bottom: 15px;
            font-size: 1.3em;
        }
        .endpoint {
            display: flex;
            align-items: center;
            padding: 10px;
            margin: 8px 0;
            background: white;
            border-radius: 8px;
            border-left: 4px solid #667eea;
        }
        .method {
            font-weight: bold;
            padding: 4px 12px;
            border-radius: 4px;
            font-size: 0.8em;
            margin-right: 15px;
            min-width: 70px;
            text-align: center;
        }
        .get { background: #61affe; color: white; }
        .post { background: #49cc90; color: white; }
        .put { background: #fca130; color: white; }
        .delete { background: #f93e3e; color: white; }
        .path {
            font-family: monospace;
            color: #333;
        }
        .links {
            display: flex;
            gap: 15px;
            justify-content: center;
            margin-top: 30px;
            flex-wrap: wrap;
        }
        .btn {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            padding: 12px 25px;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            transition: transform 0.2s, box-shadow 0.2s;
        }
        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 20px rgba(0,0,0,0.2);
        }
        .btn-primary {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
        }
        .btn-secondary {
            background: #f8f9fa;
            color: #333;
            border: 2px solid #ddd;
        }
        .architecture {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
            gap: 15px;
            margin-top: 15px;
        }
        .layer {
            text-align: center;
            padding: 15px;
            background: white;
            border-radius: 8px;
            border: 2px solid #e0e0e0;
        }
        .layer-name {
            font-weight: bold;
            color: #667eea;
        }
        .challenge {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
            color: white;
            padding: 15px;
            border-radius: 10px;
            margin-top: 20px;
        }
        .challenge h3 {
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Kasir API</h1>
        <p class="subtitle">
            <span class="badge">Week 2</span>
            <span class="badge">Layered Architecture</span>
            <span class="badge">JOIN Challenge</span>
        </p>

        <div class="section">
            <h2>📚 Layered Architecture</h2>
            <div class="architecture">
                <div class="layer">
                    <div class="layer-name">Handler</div>
                    <small>HTTP Layer</small>
                </div>
                <div class="layer">
                    <div class="layer-name">Service</div>
                    <small>Business Logic</small>
                </div>
                <div class="layer">
                    <div class="layer-name">Repository</div>
                    <small>Data Access</small>
                </div>
                <div class="layer">
                    <div class="layer-name">Entity</div>
                    <small>Models</small>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>📂 Categories Endpoints</h2>
            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/api/categories</span>
            </div>
            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/api/categories/{id}</span>
            </div>
            <div class="endpoint">
                <span class="method post">POST</span>
                <span class="path">/api/categories</span>
            </div>
            <div class="endpoint">
                <span class="method put">PUT</span>
                <span class="path">/api/categories/{id}</span>
            </div>
            <div class="endpoint">
                <span class="method delete">DELETE</span>
                <span class="path">/api/categories/{id}</span>
            </div>
        </div>

        <div class="section">
            <h2>📦 Products Endpoints</h2>
            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/api/produk</span>
            </div>
            <div class="endpoint">
                <span class="method get">GET</span>
                <span class="path">/api/produk/{id}</span>
            </div>
            <div class="endpoint">
                <span class="method post">POST</span>
                <span class="path">/api/produk</span>
            </div>
            <div class="endpoint">
                <span class="method put">PUT</span>
                <span class="path">/api/produk/{id}</span>
            </div>
            <div class="endpoint">
                <span class="method delete">DELETE</span>
                <span class="path">/api/produk/{id}</span>
            </div>
        </div>

        <div class="challenge">
            <h3>⭐ Week 2 Challenge: JOIN</h3>
            <p>Get Product Detail dengan Category Name (JOIN)</p>
            <code>GET /api/produk/1</code> akan return <code>product</code> + <code>category</code> object
        </div>

        <div class="links">
            <a href="/swagger/" class="btn btn-primary">
                📖 Swagger UI
            </a>
            <a href="/health" class="btn btn-secondary">
                💓 Health Check
            </a>
        </div>
    </div>
</body>
</html>`

func main() {
	// ===== LAYERED ARCHITECTURE SETUP =====
	
	// Repository Layer (Data Access)
	categoryRepo := repository.NewCategoryRepository()
	productRepo := repository.NewProductRepository()
	
	// Service Layer (Business Logic)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo, categoryRepo) // Inject categoryRepo untuk JOIN
	
	// Handler Layer (HTTP Handler/Controller)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)
	
	// ===== ROUTES =====
	
	// Root endpoint dengan HTML rapi
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexHTML))
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
	
	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"OK","message":"API Running with Layered Architecture"}`))
	})
	
	// Server info
	println("╔════════════════════════════════════════════════════════════╗")
	println("║                    🚀 Kasir API Server                     ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  📍 Endpoint Utama:  http://localhost:8080/                ║")
	println("║  📖 Swagger UI:      http://localhost:8080/swagger/        ║")
	println("║  💓 Health Check:    http://localhost:8080/health          ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  Layered Architecture:                                     ║")
	println("║    • Entity     → entity/                                  ║")
	println("║    • Repository → repository/                              ║")
	println("║    • Service    → service/                                 ║")
	println("║    • Handler    → handler/                                 ║")
	println("╠════════════════════════════════════════════════════════════╣")
	println("║  ⭐ Challenge: JOIN Product dengan Category!                ║")
	println("╚════════════════════════════════════════════════════════════╝")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println("❌ Gagal running server:", err.Error())
	}
}
