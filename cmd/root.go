package cmd

import (
	"fmt"
	"os"
	"snapfood/cmd/server"
	"snapfood/cmd/setup"
	"snapfood/db"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snapfood",
	Short: "A brief description of your application",
}

func Execute() {
	d := db.New()

	setup.Register(rootCmd, d)
	server.Register(rootCmd, d)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
