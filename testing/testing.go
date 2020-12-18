package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type tables struct {
	ID       int    `db:"id"`
	TypeStr  string `db:"type"`
	name     string `db:"name"`
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
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
