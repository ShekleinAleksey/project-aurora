package service

import (
	"fmt"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/ShekleinAleksey/project-aurora/internal/repository"
)

type CategoryService interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id int) (entity.Category, error)
	CreateCategory(category *entity.CreateCategoryRequest) (int, error)
	DeleteCategory(id int) error
	UpdateCategory(category *entity.Category) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) GetCategoryByID(id int) (entity.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) CreateCategory(category *entity.CreateCategoryRequest) (int, error) {
	id, err := s.repo.Create(*category)
	if err != nil {
		return id, fmt.Errorf("failed to create category: %w", err)
	}
	return id, nil
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}

func (s *categoryService) UpdateCategory(category *entity.Category) error {
	// Проверяем существование категории
	if _, err := s.repo.GetByID(category.ID); err != nil {
		return fmt.Errorf("country not found")
	}

	return s.repo.Update(*category)
}
