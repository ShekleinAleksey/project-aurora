package repository

import (
	"fmt"
	"time"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PurchaseRepository interface {
	Create(purchase entity.CreatePurchaseRequest) (int, error)
	GetAll() ([]entity.Purchase, error)
	GetByID(id int) (entity.Purchase, error)
	Update(purchase entity.Purchase) error
	Delete(id int) error
}

type purchaseRepo struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) PurchaseRepository {
	return &purchaseRepo{db: db}
}

func (r *purchaseRepo) Create(purchase entity.CreatePurchaseRequest) (int, error) {
	var id int
	now := time.Now()
	err := r.db.QueryRow("INSERT INTO purchases (name, createdat) VALUES ($1, $2) RETURNING id", now).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *purchaseRepo) GetAll() ([]entity.Purchase, error) {
	var purchases []entity.Purchase
	query := "SELECT * FROM purchases"
	err := r.db.Select(&purchases, query)
	if err != nil {
		return []entity.Purchase{}, fmt.Errorf("failed to get all purchases: %w", err)
	}
	return purchases, nil
}

func (r *purchaseRepo) GetByID(id int) (entity.Purchase, error) {
	var purchase entity.Purchase
	query := "SELECT * FROM purchases WHERE id=$1"
	err := r.db.Get(&purchase, query, id)
	if err != nil {
		return entity.Purchase{}, fmt.Errorf("failed to get purchase by id: %w", err)
	}

	return purchase, nil
}

func (r *purchaseRepo) Update(purchase entity.Purchase) error {
	query := "UPDATE purchases SET name=$1 WHERE id=$2"
	_, err := r.db.Exec(query, purchase.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *purchaseRepo) Delete(id int) error {
	query := "DELETE FROM purchases WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
