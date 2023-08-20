package cmd

import (
	"fmt"
	"gcnt/config"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cmd = &cobra.Command{
	Use:   "app",
	Short: "App GCNT",
	Long:  `Application Golang - CQRS - N.A.T.S - TypeSense`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func ExecuteCmd() {
	InitConfig()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InitConfig() {
	err := config.InitConfigInstance()
	if err != nil {
		log.Fatalf("Config error : %s", err)
	}
}
