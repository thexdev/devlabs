package app

import (
	"gohexarch/internal/domain/services"
	"gohexarch/internal/infrastructure/primary/http"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap() *fiber.App {
	// Initialize domain service
	calculatorService := services.NewCalculatorService()

	// Create HTTP handler (adapter)
	handler := http.NewCalculatorHandler(calculatorService)

	// Configure Fiber
	app := fiber.New()
	http.SetupRoutes(app, handler)

	return app
}
