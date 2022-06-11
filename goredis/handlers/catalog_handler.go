package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/goredis/services"
)

type catalogHandler struct {
	catalogService services.CatalogService
}

func NewCatalogHandler(catalogService services.CatalogService) CatalogHandler {
	return &catalogHandler{catalogService}
}

func (h *catalogHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.catalogService.GetProducts()
	if err != nil {
		c.JSON(err)
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}
	return c.JSON(response)
}
