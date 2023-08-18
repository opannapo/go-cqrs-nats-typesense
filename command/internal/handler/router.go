package handler

import (
	"gcnt/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Handler struct {
}

func (h *Handler) SetupAPIRouter() (app *fiber.App) {
	app = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  config.Instance.AppName,
		AppName:       config.Instance.AppName,
	})

	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("/api/v1")

	// User
	userGroup := api.Group("/article")
	userGroup.Post("/", func(ctx *fiber.Ctx) error {
		return nil
	})

	return
}
