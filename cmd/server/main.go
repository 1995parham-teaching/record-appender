package server

import (
	"database/sql"
	"snapfood/handler"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, db *sql.DB)  {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				api := handler.API{DB: db}
				api.Run()
			},
		},
	)
}
