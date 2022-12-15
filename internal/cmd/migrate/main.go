package migrate

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, db *sql.DB) {
	root.AddCommand(
		&cobra.Command{
			Use:   "migrate",
			Short: "Create tables",
			Run: func(cmd *cobra.Command, args []string) {
				_, err := db.Exec("CREATE TABLE IF NOT EXISTS data (" +
					"first VARCHAR NOT NULL," +
					"last VARCHAR NOT NULL," +
					"number VARCHAR NOT NULL" +
					");")
				if err != nil {
					log.Println("Cannot create data table due to the following error", err.Error())
				}
			},
		},
	)
}
