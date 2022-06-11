package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/goredis/services"
)

type catalogHandlerRedis struct {
	catalogService services.CatalogService
	redisClient    *redis.Client
}

func NewCatalogHandlerRedis(catalogService services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return &catalogHandlerRedis{catalogService, redisClient}
}

func (h *catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "handler::GetProducts"

	responseJson, err := h.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		fmt.Println("redis")
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.SendString(responseJson)
	}
	products, err := h.catalogService.GetProducts()
	if err != nil {
		c.JSON(err)
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}
	return c.JSON(response)
}
