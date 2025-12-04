package handler

import "github.com/ShekleinAleksey/project-aurora/internal/service"

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}
