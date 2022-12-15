package server

import (
	"database/sql"

	"github.com/1995parham-teaching/record-appender/internal/handler"
	"github.com/1995parham-teaching/record-appender/internal/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, db *sql.DB) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				str := store.SQLData{DB: db}
				api := handler.API{Store: str}
				api.Run()
			},
		},
	)
}
