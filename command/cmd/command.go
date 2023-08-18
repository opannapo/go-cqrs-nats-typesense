package cmd

import (
	"fmt"
	"gcnt/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	cmd.AddCommand(commandServiceCmd)
}

var commandServiceCmd = &cobra.Command{
	Use:   "command",
	Short: "Command Service",
	Long:  `Command Service`,
	Run: func(cmd *cobra.Command, args []string) {
		InitHttpServer()
	},
}

func InitHttpServer() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  config.Instance.AppName,
		AppName:       config.Instance.AppName,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading .env file : %v", err)
	}

	err := app.Listen(fmt.Sprintf(":%d", config.Instance.CommandServicePort))
	if err != nil {
		log.Panicf("Error starting server : %v", err)
	}
}
