package handlers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/gotest/handlers"
	"github.com/sing3demons/gotest/repositories"
	"github.com/sing3demons/gotest/services"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//Arrange
		amount := 100
		expected := 80
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8080/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})

	t.Run("Status BadRequest", func(t *testing.T) {
		//Arrange
		amount := "a100"
		expected := 80
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8080/calculate?amount=a100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("Status NotFound", func(t *testing.T) {
		//Arrange
		amount := 0

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(amount)
		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8080/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if err != nil {
			assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
		}
	})
}
