package app

import (
	"gohexarch/internal/domain/services"
	"gohexarch/internal/infrastructure/primary/http"
	"gohexarch/internal/infrastructure/secondary/inmemory"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap() *fiber.App {
	// Initialize driven adapter (in-memory repository)
	historyRepo := inmemory.NewHistoryRepository()

	// Create domain service with repository
	calculatorService := services.NewCalculatorService(historyRepo)

	// Inject BOTH service and repository into the handler
	handler := http.NewCalculatorHandler(calculatorService, historyRepo)

	// Configure Fiber
	app := fiber.New()
	http.SetupRoutes(app, handler)

	return app
}
