package cmd

import (
	"fmt"
	"os"

	"github.com/1995parham-teaching/record-appender/internal/cmd/migrate"
	"github.com/1995parham-teaching/record-appender/internal/cmd/server"
	"github.com/1995parham-teaching/record-appender/internal/cmd/setup"
	"github.com/1995parham-teaching/record-appender/internal/config"
	"github.com/1995parham-teaching/record-appender/internal/db"

	"github.com/spf13/cobra"
)

func Execute() {
	cfg := config.New()

	db := db.New(cfg.Database)

	rootCmd := &cobra.Command{
		Use:   "record-appender",
		Short: "read and server records",
	}

	setup.Register(rootCmd, db)
	server.Register(rootCmd, db)
	migrate.Register(rootCmd, db)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
