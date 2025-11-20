package handler

import (
	"github.com/ShekleinAleksey/project-aurora/internal/service"
)

type PurchaseHandler struct {
	service service.PurchaseService
}

func NewPurchaseHandler(service service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{service: service}
}
