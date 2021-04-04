package config

import (
	"flag"
	"fmt"
)

var (
	temp_root_dir      string = "testdata11"
	temp_Seperator     string = ","
	temp_Db_name       string = "def_database"
	temp_Table_name    string = "def_table"
	temp_Columns_names string = "-1"

	Root_dir      *string = &temp_root_dir
	Seperator     *string = &temp_Seperator
	Db_name       *string = &temp_Db_name
	Table_name    *string = &temp_Table_name
	Columns_names *string = &temp_Columns_names

	exp_root_dir      string
	exp_seperator     string
	exp_db_name       string
	exp_table_name    string
	exp_columns_names string
)

func FlagParsing() {

	//Explanation of these flags
	exp_root_dir = "Give a root directory for converting files. Example: -root=file1"
	exp_seperator = "Seperator for parsing data. Example: -seperator=,"
	exp_db_name = "Database name to be created. Example: -db=my_db"
	exp_table_name = "Columns headers to be created. Example: -columns=col1,col2,col3"
	exp_columns_names = "Table name to be created. Example: -table=my_table"

	//Getting flags from CLI
	Root_dir = flag.String("root", *Root_dir, exp_root_dir)
	Seperator = flag.String("seperator", *Seperator, exp_seperator)
	Db_name = flag.String("db", *Db_name, exp_db_name)
	Columns_names = flag.String("columns", *Columns_names, exp_columns_names)
	Table_name = flag.String("table", *Table_name, exp_table_name)

	flag.Parse()

	fmt.Println("Root Directory: ", *Root_dir)
	fmt.Println("Seperator: ", *Seperator)
	fmt.Println("Database Name: ", *Db_name)
	fmt.Println("Column Names: ", *Columns_names)
	fmt.Println("Table Name: ", *Table_name)
	fmt.Println("----------------------------------")

}
