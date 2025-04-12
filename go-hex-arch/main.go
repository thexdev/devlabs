package main

import (
	"gohexarch/internal/app/usecases"
	"gohexarch/internal/infra/primary/http"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		addUseCase      = new(usecases.AddUseCase)
		subtractUseCase = new(usecases.SubstractUseCase)
		devideUseCase   = new(usecases.DevideUseCase)
		multiplyUseCase = new(usecases.MultiplyUseCase)
	)

	router := fiber.New()

	handler := http.NewCalculatorHandler(
		addUseCase,
		subtractUseCase,
		devideUseCase,
		multiplyUseCase,
	)

	router.Post("/compute", handler.Compute)

	log.Fatal(router.Listen(":3000"))
}
