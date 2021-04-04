package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Yavuzlar/CSVtoSQLite/config"
	"github.com/Yavuzlar/CSVtoSQLite/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Getting flags
	config.FlagParsing()

	// Initialize Database
	database.InitDatabase()

	err := filepath.Walk("./"+*config.Root_dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					log.Fatalf("failed to open")
				}
				readerCSV := csv.NewReader(file)
				readerCSV.Comma = ','

				ctx := context.Background()
				tx, err := database.Db.BeginTx(ctx, nil)
				if err != nil {
					log.Fatal(err)
				}

				for i := 1; ; i = i + 1 {
					record, err := readerCSV.Read()
					if err == io.EOF {
						break
					} else if err != nil {
						fmt.Println("An error encountered ::", err)

					}
					s_interface := make([]interface{}, len(record))
					for i, v := range record {
						s_interface[i] = v
					}

					_, err = tx.ExecContext(ctx, database.InsertQuery, s_interface...)
					if err != nil {
						fmt.Println(err)
						tx.Rollback()
						break
					}
					fmt.Println(i)
					if i > 100000 {
						break
					}
				}

				err = tx.Commit()
				if err != nil {
					log.Fatal(err)
				}

				file.Close()

				return nil
			}
			return nil

		})
	if err != nil {
		log.Println(err)
	}

}
