package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	PurchaseRepository PurchaseRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PurchaseRepository: NewPurchaseRepository(db),
	}
}
