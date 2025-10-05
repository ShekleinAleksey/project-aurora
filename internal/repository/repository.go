package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ExpenseRepository *ExpenseRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ExpenseRepository: NewExpenseRepository(db),
	}
}
