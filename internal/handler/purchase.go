package handler

import "github.com/ShekleinAleksey/project-aurora/internal/entity"

type ExpenseService interface {
	GetAllExpansies() ([]entity.Purchase, error)
	GetExpanseByID(id int) (entity.Purchase, error)
}

type ExpenseHandler struct {
	ExpenseService ExpenseService
}

func NewExpenseHandler(expenseService ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{ExpenseService: expenseService}
}
