package service

import "github.com/ShekleinAleksey/project_Aurora/internal/repository"

type Deps struct {
	ExpenseRepository *repository.ExpenseRepository
}

type Service struct {
	ExpenseService *ExpenseService
}

func NewService(deps Deps) *Service {
	return &Service{
		ExpenseService: NewExpenseService(*deps.ExpenseRepository),
	}
}
