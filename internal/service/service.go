package service

import "github.com/ShekleinAleksey/project-aurora/internal/repository"

type Service struct {
	CategoryService CategoryService
	PurchaseService PurchaseService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		CategoryService: NewCategoryService(r.CategoryRepository),
		PurchaseService: NewPurchaseService(r.PurchaseRepository),
	}
}
