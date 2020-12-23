package main

import (
	"database/sql"
	"fmt"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"github.com/jmoiron/sqlx"
	"github.com/leaanthony/mewn"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails"
)

var (
	db *sqlx.DB
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{

		Title:     "GoLite",
		JS:        js,
		CSS:       css,
		Colour:    "#1c1e26",
		Resizable: true,
	})
	app.Bind(selectDatabase)
	app.Bind(selectTable)
	app.Bind(getHeaders)
	app.Run()
}

func selectDatabase() []tables {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open a database file",
		Role:  "GoLite-UI",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "SQLite-files (*.db, *.db3, *.sqlite, *.sqlite3)",
				Pattern:     "*.db;*.db3;*.sqlite;*.sqlite3",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	db = connectDatabase(result)
	return getTables(db)
}

type headers struct {
	Cid       string         `db:"cid"`
	Name      string         `db:"name"`
	Type      string         `db:"type"`
	NotNull   int            `db:"notnull"`
	DfltValue sql.NullString `db:"dflt_value"`
	Pk        int            `db:"pk"`
}

func getHeaders(tableName string) []headers {
	results := []headers{}
	query := "PRAGMA table_info(" + tableName + ")"

	errQuery := db.Select(&results, query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	return results
}

func selectTable(tableName string) []string {
	results := []string{}
	errQuery := db.Select(&results, "SELECT * FROM ?", tableName)
	if errQuery != nil {
		fmt.Println(errQuery)
	}
	fmt.Println("#############################")
	fmt.Println(tableName)
	fmt.Println(results)
	return results
}

func connectDatabase(path string) *sqlx.DB {
	dbConnection, err := sqlx.Connect("sqlite3", path)

	if err != nil {
		fmt.Println(err)
	}
	return dbConnection
}

type tables struct {
	TypeStr  string `db:"type"`
	Name     string `db:"name"`
	TblName  string `db:"tbl_name"`
	Rootpage int    `db:"rootpage"`
	SQL      string `db:"sql"`
}

func getTables(connection *sqlx.DB) []tables {
	results := []tables{}
	errQuery := connection.Select(&results, "SELECT * FROM sqlite_master where type='table'")
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	return results
}
