package handler

import "github.com/ShekleinAleksey/project_Aurora/internal/entity"

type ExpenseService interface {
	GetAllExpansies() ([]entity.Expense, error)
	GetExpanseByID(id int) (entity.Expense, error)
}

type ExpenseHandler struct {
	ExpenseService ExpenseService
}

func NewExpenseHandler(expenseService ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{ExpenseService: expenseService}
}
