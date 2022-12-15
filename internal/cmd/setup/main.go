package setup

import (
	"bufio"
	"database/sql"
	"errors"
	"io"
	"log"
	"os"

	"github.com/1995parham-teaching/record-appender/internal/model"
	"github.com/1995parham-teaching/record-appender/internal/store"

	"github.com/spf13/cobra"
)

const BulkRead = 2

func Register(root *cobra.Command, db *sql.DB) {
	root.AddCommand(
		&cobra.Command{
			Use:   "setup",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				f, err := os.Open("./data.txt")
				if err != nil {
					log.Fatal(err)
				}

				str := store.SQLData{DB: db}

				r := bufio.NewReader(f)

				for {
					models := make([]model.Data, 0)
					var err error
					for i := 0; i < BulkRead; i++ {
						var line string
						line, err = r.ReadString('\n')
						if err != nil {
							if errors.Is(err, io.EOF) {
								break
							}

							log.Println(err)
							return
						}

						models = append(models, model.New(line))
					}

					if err := str.Insert(models); err != nil {
						log.Println(err)
					}

					if errors.Is(err, io.EOF) {
						break
					}
				}
			},
		},
	)
}
