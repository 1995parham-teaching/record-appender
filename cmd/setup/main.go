package setup

import (
	"bufio"
	"database/sql"
	"errors"
	"io"
	"log"
	"os"
	"snapfood/model"
	"snapfood/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *sql.DB) {
	root.AddCommand(
		&cobra.Command{
			Use:   "setup",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				f, err := os.Open("./data.txt")
				check(err)

				r := bufio.NewReader(f)

				for {
					models := make([]model.Data, 0)
					var err error
					for i := 0; i < 2; i++ {
						var line string
						line, err = r.ReadString('\n')
						if err != nil {
							if errors.Is(err, io.EOF) {
								break
							}

							log.Println(err)
						}
						models = append(models, model.New(line))
					}

					if err := store.Insert(d, models); err != nil {
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

func check(err error)  {
	if err != nil {
		log.Println(err)
	}
}