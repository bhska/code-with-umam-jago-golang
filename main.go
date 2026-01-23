package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	CategoryID int    `json:"category_id"`
}

var categories = []Category{
	{ID: 1, Name: "Minuman", Description: "Segala jenis minuman"},
	{ID: 2, Name: "Makanan", Description: "Segala jenis makanan"},
}

var produk = []Produk{
	{ID: 1, Nama: "Produk A", Harga: 10000, CategoryID: 1},
	{ID: 2, Nama: "Produk B", Harga: 20000, CategoryID: 2},
	{ID: 3, Nama: "Produk C", Harga: 30000, CategoryID: 1},
}

func ambilSemuaKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func tambahKategori(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru Category
	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(categories) + 1
	categories = append(categories, kategoriBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriBaru)
}

func ambilKategoriByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func updateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateKategori Category
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			updateKategori.ID = id
			categories[i] = updateKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func hapusKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Category deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func ambilProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func hapusProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func main() {
	// Category endpoints
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ambilKategoriByID(w, r)
		case "PUT":
			updateKategori(w, r)
		case "DELETE":
			hapusKategori(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ambilSemuaKategori(w, r)
		case "POST":
			tambahKategori(w, r)
		}
	})

	// Produk endpoints
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			ambilProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			hapusProduk(w, r)
		}
	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
