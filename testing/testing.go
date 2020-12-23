package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type tables struct {
	// ID       int    `db:"id"`
	TypeStr  string `db:"type"`
	Name     string `db:"name"`
	TblName  string `db:"tbl_name"`
	Rootpage int    `db:"rootpage"`
	SQL      string `db:"sql"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "chinook.db")
	checkErr(err)

	type headers struct {
		Cid       string         `db:"cid"`
		Name      string         `db:"name"`
		Type      string         `db:"type"`
		NotNull   int            `db:"notnull"`
		DfltValue sql.NullString `db:"dflt_value"`
		Pk        int            `db:"pk"`
	}

	results := []headers{}
	tableName := "albums"
	query := "PRAGMA table_info(" + tableName + ")"

	errQuery := db.Select(&results, query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	fmt.Println(results)

	// test := "albums"
	// query := "SELECT * FROM " + test
	// results, errQuery := db.Query(query)
	// if errQuery != nil {
	// 	fmt.Println(errQuery)
	// }
	// fmt.Println("#############################")
	// fmt.Println(results)
	// fmt.Println("teeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeest")
	// result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
	// 	Title: "Open A File",
	// 	Role:  "OpenFileExample",
	// 	FileFilters: []cfd.FileFilter{
	// 		{
	// 			DisplayName: "Text Files (*.txt)",
	// 			Pattern:     "*.txt",
	// 		},
	// 		{
	// 			DisplayName: "Image Files (*.jpg, *.png)",
	// 			Pattern:     "*.jpg;*.png",
	// 		},
	// 		{
	// 			DisplayName: "All Files (*.*)",
	// 			Pattern:     "*.*",
	// 		},
	// 	},
	// 	SelectedFileFilterIndex: 2,
	// 	FileName:                "file.txt",
	// 	DefaultExtension:        "txt",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result, "#########################")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
