package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/ShekleinAleksey/project-aurora/internal/repository"
)

type PurchaseService interface {
	GetAllPurchases() ([]entity.Purchase, error)
	GetPurchaseByID(id int) (entity.Purchase, error)
	CreatePurchase(purchase *entity.CreatePurchaseRequest) (int, error)
	DeletePurchase(id int) error
	UpdatePurchase(purchase *entity.Purchase) error
	SumTotalPrice() (float64, error)
	BatchPurchases(purchases []entity.CreatePurchaseRequest) error
	ImportFromJSON(filePath string) (int, error)
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

func (s *purchaseService) SumTotalPrice() (float64, error) {
	return s.repo.SumTotalPrice()
}

func (s *purchaseService) BatchPurchases(purchases []entity.CreatePurchaseRequest) error {
	for _, purchase := range purchases {
		if _, err := s.CreatePurchase(&purchase); err != nil {
			return fmt.Errorf("failed to create purchase: %w", err)
		}
	}
	return nil
}

func (s *purchaseService) ImportFromJSON(filePath string) (int, error) {
	// Читаем файл
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Парсим JSON
	var purchases []entity.CreatePurchaseRequest
	err = json.Unmarshal(data, &purchases)
	if err != nil {
		return 0, fmt.Errorf("failed to parse JSON: %w", err)
	}

	err = s.BatchPurchases(purchases)
	if err != nil {
		return 0, fmt.Errorf("failed to batch create purchases: %w", err)
	}

	return len(purchases), nil
}
