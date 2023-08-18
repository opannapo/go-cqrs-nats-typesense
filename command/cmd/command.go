package cmd

import (
	"fmt"
	"gcnt/config"
	"gcnt/internal/handler"
	"gcnt/internal/repository"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(commandServiceCmd)
}

var commandServiceCmd = &cobra.Command{
	Use:   "command",
	Short: "Command Service",
	Long:  `Command Service`,
	Run: func(cmd *cobra.Command, args []string) {
		InitDatabase()
		InitHttpServer()
	},
}

func InitHttpServer() {
	app := handler.SetupAPIRouter()
	err := app.Listen(fmt.Sprintf(":%d", config.Instance.CommandServicePort))
	if err != nil {
		log.Panic().Caller().Err(err)
	}
}

func InitDatabase() {
	db := &repository.Db{}
	err := db.InitDatabase("mysql")
	sqlDB, err := db.Mysql.DB()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
		return
	}
	defer sqlDB.Close()
}
