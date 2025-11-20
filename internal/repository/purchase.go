package repository

import "github.com/jmoiron/sqlx"

type PurchaseRepository interface {
}

type purchaseRepo struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) PurchaseRepository {
	return &purchaseRepo{db: db}
}
