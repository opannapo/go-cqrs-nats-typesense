package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{
	Use:   "app",
	Short: "App GCNT",
	Long:  `Application Golang - CQRS - N.A.T.S - TypeSense`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func ExecuteCmd() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
