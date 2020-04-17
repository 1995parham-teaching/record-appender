package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/record-appender/cmd/migrate"
	"github.com/elahe-dastan/record-appender/cmd/server"
	"github.com/elahe-dastan/record-appender/cmd/setup"
	"github.com/elahe-dastan/record-appender/config"
	"github.com/elahe-dastan/record-appender/db"

	"github.com/spf13/cobra"
)

func Execute() {
	cfg := config.New()

	db := db.New(cfg.Database)

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "record-appender",
		Short: "A brief description of your application",
	}

	setup.Register(rootCmd, db)
	server.Register(rootCmd, db)
	migrate.Register(rootCmd, db)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
