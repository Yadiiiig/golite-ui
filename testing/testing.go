package main

import (
	"fmt"
	"log"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
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
	db, err := sqlx.Connect("sqlite3", "bugs.db")
	checkErr(err)

	results := []tables{}
	errQuery := db.Select(&results, "SELECT * FROM sqlite_master where type='table'")
	checkErr(errQuery)

	fmt.Println(results)
	fmt.Println("teeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeest")
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.txt",
		DefaultExtension:        "txt",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result, "#########################")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
