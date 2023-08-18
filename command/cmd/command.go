package cmd

import (
	"fmt"
	"gcnt/config"
	"gcnt/internal/handler"
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
	h := handler.Handler{}
	app := h.SetupAPIRouter()

	err := app.Listen(fmt.Sprintf(":%d", config.Instance.CommandServicePort))
	if err != nil {
		log.Panicf("Error starting server : %v", err)
	}
}
