package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	CategoryRepository CategoryRepository
	PurchaseRepository PurchaseRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CategoryRepository: NewCategoryRepository(db),
		PurchaseRepository: NewPurchaseRepository(db),
	}
}
