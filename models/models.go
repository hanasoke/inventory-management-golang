package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
	MinStock    int       `gorm:"default:5" json:"min_stock"`
	SKU         string    `gorm:"uniqueIndex" json:"sku"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Supplier struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `gorm:"not null" json:"name"`
    Contact     string    `json:"contact"`
    Phone       string    `json:"phone"`
    Email       string    `json:"email"`
    Address     string    `json:"address"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    ProductID   uint      `gorm:"not null" json:"product_id"`
    Product     Product   `gorm:"foreignKey:ProductID" json:"product"`
    Type        string    `gorm:"not null" json:"type"` // "in" or "out"
    Quantity    int       `gorm:"not null" json:"quantity"`
    Price       float64   `json:"price"`
    Notes       string    `json:"notes"`
    CreatedAt   time.Time `json:"created_at"`
}

type StockAlert struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    ProductID   uint      `gorm:"not null" json:"product_id"`
    Product     Product   `gorm:"foreignKey:ProductID" json:"product"`
    Message     string    `gorm:"not null" json:"message"`
    IsResolved  bool      `gorm:"default:false" json:"is_resolved"`
    CreatedAt   time.Time `json:"created_at"`
}