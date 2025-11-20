package service

import "github.com/ShekleinAleksey/project-aurora/internal/repository"

type PurchaseService interface {
}

type purchaseService struct {
	repo repository.PurchaseRepository
}

func NewPurchaseService(repo repository.PurchaseRepository) PurchaseService {
	return &purchaseService{repo: repo}
}
