package main

import (
	"gohexarch/internal/app/usecases"
	"gohexarch/internal/infra/primary/http"
	"gohexarch/internal/infra/secondary/repository"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	var (
		repo = repository.NewInMemoryRespository()

		addUseCase      = usecases.NewAddUseCase(repo)
		subtractUseCase = usecases.NewSubtractUseCase(repo)
		devideUseCase   = usecases.NewDevideUseCase(repo)
		multiplyUseCase = usecases.NewMultiplyUseCase(repo)
	)

	router := fiber.New()
	
	router.Use(recover.New())

	calcHandler := http.NewCalculatorHandler(
		addUseCase,
		subtractUseCase,
		devideUseCase,
		multiplyUseCase,
	)

	historyHandler := http.NewHistoryHandler(repo)

	router.Post("/compute", calcHandler.Compute)
	router.Get("/history", historyHandler.All)

	log.Fatal(router.Listen(":3000"))
}
