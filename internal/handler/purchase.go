package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	service service.PurchaseService
}

func NewPurchaseHandler(service service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{service: service}
}

// CreatePurchase создает новую покупку
// @Summary Создать покупку
// @Description Создает новую запись о покупке
// @Tags purchases
// @Accept json
// @Produce json
// @Param request body string true "Данные покупки"
// @Success 201 {object} entity.Purchase
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchases [post]
func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {

	var purchase entity.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreatePurchase(&purchase)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetAllPurchases возвращает список покупок
// @Summary Список покупок
// @Description Возвращает список всех покупок
// @Tags purchases
// @Produce json
// @Success 200 {array} entity.Purchase
// @Failure 500 {object} map[string]string
// @Router /purchases [get]
func (h *PurchaseHandler) GetAllPurchases(c *gin.Context) {
	purchases, err := h.service.GetAllPurchases()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, purchases)
}

// GetPurchase возвращает покупку по ID
// @Summary Получить покупку
// @Description Возвращает покупку по ID
// @Tags purchases
// @Produce json
// @Param id path string true "ID покупки"
// @Success 200 {array} entity.Purchase
// @Failure 500 {object} map[string]string
// @Router /purchases/{id} [get]
func (h *PurchaseHandler) GetPurchase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	purchase, err := h.service.GetPurchaseByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, purchase)
}

// UpdatePurchases обновляет покупку
// @Summary Обновить покупку
// @Description Обновляет данные покупки по ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "ID покупки"
// @Param request body string true "Данные для обновления"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchases/{id} [put]
func (h *PurchaseHandler) UpdatePurchases(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid purchase ID"})
		return
	}

	var purchase entity.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	purchase.ID = id

	err = h.service.UpdatePurchase(&purchase)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "purchase not found" {
			status = http.StatusNotFound
		} else if strings.Contains(err.Error(), "is required") {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// DeletePurchase удаляет покупку
// @Summary Удалить покупку
// @Description Удаляет покупку по ID
// @Tags purchases
// @Produce json
// @Param id path string true "ID покупки"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchases/{id} [delete]
func (h *PurchaseHandler) DeletePurchase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.service.DeletePurchase(id)
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
