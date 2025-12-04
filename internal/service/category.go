package service

import "github.com/ShekleinAleksey/project-aurora/internal/repository"

type CategoryService interface {
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}
