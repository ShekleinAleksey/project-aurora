package entity

import "time"

type Purchase struct {
	ID           int       `json:"id" db:"id"`
	Material     string    `json:"material" db:"material"`
	Count        float64   `json:"count" db:"count"`
	UnitPrice    float64   `json:"unit_price" db:"unit_price"`
	TotalPrice   float64   `json:"total_price" db:"total_price"`
	Notes        string    `json:"notes" db:"notes"`
	PurchaseDate time.Time `json:"purchase_date" db:"purchase_date"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreatePurchaseRequest struct {
	Material     string    `json:"material" binding:"required"`
	Count        float64   `json:"count" binding:"required,gt=0"`
	UnitPrice    float64   `json:"unit_price" binding:"required,gt=0"`
	Notes        string    `json:"notes"`
	PurchaseDate time.Time `json:"purchase_date" binding:"required"`
}
