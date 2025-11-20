package service

import "github.com/ShekleinAleksey/project-aurora/internal/repository"

type Service struct {
	PurchaseService PurchaseService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		PurchaseService: NewPurchaseService(r.PurchaseRepository),
	}
}
