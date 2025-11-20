# Inventory Management System

Sistem manajemen inventory toko yang dibangun dengan Golang dan Gin framework

## Fitur

- Manajemen Produk (CRUD)
- Manajemen Supplier 
- Transaksi Stok (Masuk/Keluar)
- Alert Stock Rendah
- Dashboard Statistik
- SQLite Database

## Instalasi 

1. Clone Repository 
2. Jalankan `go mod tidy` untuk menginstall dependencies
3. Jalankan `go run main.go` untuk memulai server 

## API Endpoints 

### Products 
- `Get /api/products` - Get all products 
- `GET /api/products/:id` - Get product by ID 
- `POST /api/products` - Create product baru 
- `PUT /api/products/:id` - Update product 
- `DELETE /api/products/:id` - Delete product 