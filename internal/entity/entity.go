package entity

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Material struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Category      Category  `json:"category"`
	Description   string    `json:"description"`
	ArticleNumber string    `json:"article_number"`
	Unit          string    `json:"unit"`
	Remains       int       `json:"remains"`
	MinCount      int       `json:"min_count"`
	CreatedAt     time.Time `json:"created_at"`
}

type Purchase struct {
	ID           int       `json:"id"`
	MaterialID   int       `json:"material_id"`
	Count        float64   `json:"count"`
	UnitPrice    float64   `json:"unit_price"`
	TotalPrice   float64   `json:"total_price"`
	Notes        string    `json:"notes"`
	PurchaseDate time.Time `json:"purchase_date"`
	CreatedAt    time.Time `json:"created_at"`
}
