package handler

import (
	"gcnt/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAPIRouter() (app *fiber.App) {
	ah := NewArticleHandler()
	app = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  config.Instance.AppName,
		AppName:       config.Instance.AppName,
	})
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api := app.Group("/api/v1")
	userGroup := api.Group("/article")
	userGroup.Post("/", ah.CreateArticle)

	return
}
