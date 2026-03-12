package entity

import "time"

type Purchase struct {
	ID           int       `json:"id"`
	Material     string    `json:"material"`
	Count        float64   `json:"count"`
	UnitPrice    float64   `json:"unit_price"`
	TotalPrice   float64   `json:"total_price"`
	Notes        string    `json:"notes"`
	PurchaseDate time.Time `json:"purchase_date"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreatePurchaseRequest struct {
	Material     string    `json:"material" binding:"required"`
	Count        float64   `json:"count" binding:"required,gt=0"`
	UnitPrice    float64   `json:"unit_price" binding:"required,gt=0"`
	Notes        string    `json:"notes"`
	PurchaseDate time.Time `json:"purchase_date" binding:"required"`
}
