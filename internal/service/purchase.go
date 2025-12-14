package service

import (
	"fmt"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/ShekleinAleksey/project-aurora/internal/repository"
)

type PurchaseService interface {
	GetAllPurchases() ([]entity.Purchase, error)
	GetPurchaseByID(id int) (entity.Purchase, error)
	CreatePurchase(purchase *entity.CreatePurchaseRequest) (int, error)
	DeletePurchase(id int) error
	UpdatePurchase(purchase *entity.Purchase) error
}

type purchaseService struct {
	repo repository.PurchaseRepository
}

func NewPurchaseService(repo repository.PurchaseRepository) PurchaseService {
	return &purchaseService{repo: repo}
}

func (s *purchaseService) GetAllPurchases() ([]entity.Purchase, error) {
	return s.repo.GetAll()
}

func (s *purchaseService) GetPurchaseByID(id int) (entity.Purchase, error) {
	return s.repo.GetByID(id)
}

func (s *purchaseService) CreatePurchase(purchase *entity.CreatePurchaseRequest) (int, error) {
	id, err := s.repo.Create(*purchase)
	if err != nil {
		return id, fmt.Errorf("failed to create purchase: %w", err)
	}
	return id, nil
}

func (s *purchaseService) DeletePurchase(id int) error {
	return s.repo.Delete(id)
}

func (s *purchaseService) UpdatePurchase(purchase *entity.Purchase) error {
	// Проверяем существование покупки
	if _, err := s.repo.GetByID(purchase.ID); err != nil {
		return fmt.Errorf("purchase not found")
	}

	return s.repo.Update(*purchase)
}
