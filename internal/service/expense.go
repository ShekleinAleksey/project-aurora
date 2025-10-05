package service

import "github.com/ShekleinAleksey/project_Aurora/internal/repository"

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}
