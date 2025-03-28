package main

import (
	"log"
	"scalt/internal/app/service"
	"scalt/internal/infra/adapter/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	calculatorService := service.NewCalculatorService()

	handler := http.NewCalculatorHandler(calculatorService)

	app.Post("/calculate", handler.Calculate)

	log.Fatal(app.Listen(":3000"))
}
