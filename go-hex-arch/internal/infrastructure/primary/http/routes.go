package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, handler *CalculatorHandler) {
	app.Post("/calculate", handler.Calculate)
	app.Get("/history", handler.GetHistory)
}
