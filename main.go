package main

import (
	"database/sql"
	"encoding/json"
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
	app.Bind(runQuery)
	app.Bind(getTableInfo)
	app.Bind(closeDatabase)
	app.Run()
}

func runQuery(query string) interface{} {
	type queryStruct struct {
		Columns []string
		Data    string
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	result := make([]map[string]interface{}, 0)
	data := make([]interface{}, len(columns))
	for rows.Next() {
		colData := make(map[string]interface{}, len(columns))
		for i := range data {
			data[i] = new(interface{})
		}
		if err := rows.Scan(data...); err != nil {
			fmt.Println(err)
		}
		for i, col := range columns {
			colData[col] = *data[i].(*interface{})
		}
		result = append(result, colData)
	}

	finalResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	results := queryStruct{columns, string(finalResult)}

	return results
}

// type selectDatabaseResult struct {
// 	Tables []tables
// 	Info []tableInfo
// }

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
	// returnResult := []selectDatabaseResult{getTables(db)}

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

func selectTable(tableName string) interface{} {
	type queryStruct struct {
		Columns []string
		Data    string
	}

	query := "SELECT * FROM " + tableName
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	result := make([]map[string]interface{}, 0)
	data := make([]interface{}, len(columns))
	for rows.Next() {
		colData := make(map[string]interface{}, len(columns))
		for i := range data {
			data[i] = new(interface{})
		}
		if err := rows.Scan(data...); err != nil {
			fmt.Println(err)
		}
		for i, col := range columns {
			colData[col] = *data[i].(*interface{})
		}
		result = append(result, colData)
	}

	finalResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	results := queryStruct{columns, string(finalResult)}

	return results
}

func connectDatabase(path string) *sqlx.DB {
	var dbCheck *sqlx.DB
	if dbCheck != db {
		db.Close()
	}
	dbConnection, err := sqlx.Connect("sqlite3", path)

	if err != nil {
		fmt.Println(err)
	}
	return dbConnection
}

func closeDatabase() {
	var dbCheck *sqlx.DB
	if dbCheck != db {
		db.Close()
	}
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

type tableInfo struct {
	Cid       int            `db:"cid"`
	Name      string         `db:"name"`
	Type      string         `db:"type"`
	NotNull   bool           `db:"notnull"`
	DfltValue sql.NullString `db:"dflt_value"`
	Pk        int            `db:"pk"`
}

func getTableInfo(tableName string) []tableInfo {
	//fmt.Println(tableName)
	results := []tableInfo{}
	query := "PRAGMA table_info(" + tableName + ")"

	errQuery := db.Select(&results, query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}
	//fmt.Println(results)
	return results
}
