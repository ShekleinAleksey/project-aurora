package entity

import "time"

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"createdat"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
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
