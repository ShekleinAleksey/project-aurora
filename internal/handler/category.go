package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// CreateCategory создает новую категорию
// @Summary Создать категорию
// @Description Создает новую запись о категории
// @Tags categories
// @Accept json
// @Produce json
// @Param request body string true "Данные категории"
// @Success 201 {object} entity.Category
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	var category entity.CreateCategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateCategory(&category)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetAllCategories возвращает список категорий
// @Summary Список категорий
// @Description Возвращает список всех категорий
// @Tags categories
// @Produce json
// @Success 200 {array} entity.Category
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetAllCategories возвращает категорию по ID
// @Summary Получить категорию
// @Description Возвращает категорию по ID
// @Tags categories
// @Produce json
// @Param id path string true "ID подписки"
// @Success 200 {array} entity.Category
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateSubscription обновляет категорию
// @Summary Обновить категорию
// @Description Обновляет данные категории по ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "ID категории"
// @Param request body string true "Данные для обновления"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	category.ID = id

	err = h.service.UpdateCategory(&category)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "category not found" {
			status = http.StatusNotFound
		} else if strings.Contains(err.Error(), "is required") {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// DeleteSubscription удаляет категорию
// @Summary Удалить категорию
// @Description Удаляет категорию по ID
// @Tags categories
// @Produce json
// @Param id path string true "ID категории"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.service.DeleteCategory(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
