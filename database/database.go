package database

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yavuzlar/CSVtoSQLite/config"
	_ "github.com/mattn/go-sqlite3"
)

var InsertQuery string         //
var Arr_columns_names []string // Column isimleri array olarak burada bulunur
//*config.Columns_names içerisinde sütun isimleri string olarak virgüllerle ayrılmış şekilde tutulur
var Db *sql.DB

func initColumns() {

	if *config.Columns_names != "-1" {
		//Column isimleri parametre olarak verilirse direk onlar kullanılır
		Arr_columns_names = strings.Split(*config.Columns_names, ",")

	} else if *config.Columns_names == "-1" {
		//Eğer column isimleri parametre olarak verilmez ise ilk sütun verilerini column names olarak alırız

		var files []string
		err := filepath.Walk("./"+*config.Root_dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, path)
				return nil
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		csvfile, err := os.Open(files[0])
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		csvReader := csv.NewReader(csvfile)

		csvReader.Comma = ';'
		tempArr, _ := csvReader.Read()

		// Headerlarda boş(yanlış) veri varsa onu eklemeyecek
		for _, element := range tempArr {
			if element != "" {
				Arr_columns_names = append(Arr_columns_names, strings.TrimSpace(element))
			}
		}
		*config.Columns_names = strings.Join(Arr_columns_names, ",")
		csvfile.Close()
	}

	InsertQuery = "INSERT INTO " + *config.Table_name + " (" + *config.Columns_names + ") VALUES ("
	for i := 0; i < len(Arr_columns_names); i++ {
		InsertQuery += "?,"
	}
	InsertQuery = InsertQuery[:len(InsertQuery)-1] + ");"
}
func InitDatabase() {
	initColumns()

	Db, _ = sql.Open("sqlite3", "./"+*config.Db_name+".db")
	var createTableSQL string = `CREATE TABLE ` + *config.Table_name + ` (`
	for i := 0; i < len(Arr_columns_names); i++ {
		createTableSQL += `"` + Arr_columns_names[i] + `" TEXT,`
	}
	createTableSQL = createTableSQL[:len(createTableSQL)-1]
	createTableSQL += `);`

	log.Println(createTableSQL)
	log.Println("CREATING " + *config.Table_name + " TABLE...")
	statement, err := Db.Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println(*config.Table_name + " TABLE CREATED!")

}
