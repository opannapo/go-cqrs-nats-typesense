package cmd

import (
	"fmt"
	"gcnt/config"
	"gcnt/internal/handler"
	"gcnt/internal/repository"
	"gcnt/internal/service"
	"gcnt/internal/stream"
	"gcnt/internal/stream/publisher"
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
		BootStrap()
	},
}

func BootStrap() {
	//Init Service
	service.InitArticleServiceInstance()

	//Init Repo
	repository.InitArticleRepositoryInstance()

	//Init message broker
	err := stream.InitMessageBrokerInstance()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
		return
	}

	//Init publisher
	publisher.InitNatsPublisherInstance()

	//Init Database
	db := &repository.Db{}
	err = db.InitDatabaseInstance("mysql")
	sqlDB, err := db.Mysql.DB()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
		return
	}
	defer sqlDB.Close()

	//Init & Start HTTP server
	app := handler.SetupAPIRouter()
	err = app.Listen(fmt.Sprintf(":%d", config.Instance.CommandServicePort))
	if err != nil {
		log.Panic().Caller().Err(err)
	}
}
