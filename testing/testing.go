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

	type tableInfo struct {
		Cid       int            `db:"cid"`
		Name      string         `db:"name"`
		Type      string         `db:"type"`
		NotNull   bool           `db:"notnull"`
		DfltValue sql.NullString `db:"dflt_value"`
		Pk        bool           `db:"pk"`
	}

	results := []tableInfo{}
	query := "PRAGMA table_info(albums)"

	errQuery := db.Select(&results, query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	fmt.Println(results)

	// results := []headers{}
	// tableName := "albums"
	// query := "PRAGMA table_info(" + tableName + ")"

	// errQuery := db.Select(&results, query)
	// if errQuery != nil {
	// 	fmt.Println(errQuery)
	// }

	// fmt.Println(results)
	// query := "SELECT * FROM albums"
	// rows, err := db.Query(query)
	// if err != nil {
	// 	panic(err)
	// }
	// columns, err := rows.Columns()
	// if err != nil {
	// 	panic(err)
	// }

	// result := make([]map[string]interface{}, 0)
	// data := make([]interface{}, len(columns))
	// for rows.Next() {
	// 	colData := make(map[string]interface{}, len(columns))
	// 	for i := range data {
	// 		data[i] = new(interface{})
	// 	}
	// 	if err := rows.Scan(data...); err != nil {
	// 		panic(err)
	// 	}
	// 	for i, col := range columns {
	// 		colData[col] = *data[i].(*interface{})
	// 	}
	// 	result = append(result, colData)
	// }

	// //fmt.Println(result)
	// finalResult, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(finalResult))

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
