package bootstrap

import (
	"gcnt/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

var commandServiceCmd = &cobra.Command{
	Use:   "command",
	Short: "Command Service",
	Long:  `Command Service`,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		InitHttpServer()
	},
}

func init() {
	cmd.AddCommand(commandServiceCmd)
}

func InitConfig() {
	err := config.InitConfigInstance()
	if err != nil {
		log.Fatalf("Config error : %s", err)
	}
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
		log.Fatal("Error loading .env file")
	}
}
