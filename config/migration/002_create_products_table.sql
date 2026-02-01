-- Migration: Create products table
-- Created at: 2026-02-01

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    harga INTEGER NOT NULL CHECK (harga >= 0),
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for faster lookup
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_nama ON products(nama);
