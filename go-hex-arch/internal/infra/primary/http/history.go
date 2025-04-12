package http

import (
	"gohexarch/internal/app/ports/secondary"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	repo secondary.Repository
}

func NewHistoryHandler(repo secondary.Repository) *HistoryHandler {
	return &HistoryHandler{
		repo: repo,
	}
}

func (h *HistoryHandler) All(c *fiber.Ctx) error {
	return c.JSON(h.repo.All())
}
