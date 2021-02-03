package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	//"github.com/jmoiron/sqlx"
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
	// db, err := sqlx.Connect("sqlite3", "chinook.db")
	// checkErr(err)

	type createTableStruct struct {
		Name          string `json:"name"`
		Type          string `json:"type"`
		PrimaryKey    bool   `json:"PK"`
		NotNull       bool   `json:"NN"`
		AutoIncrement bool   `json:"AI"`
		Unique        bool   `json:"UN"`
	}

	type data struct {
		TableName string              `json:"tableName"`
		AmountPk  int                 `json:"apk"`
		Tables    []createTableStruct `json:"tables"`
	}

	tableData := `[{"apk": 1, "tableName": "testing", "tables": [{"name": "ArtistId", "type": "INTEGER", "NN": true, "PK": true, "AI": true, "UN": true}, {"name": "ArtistName", "type": "STRING", "NN": true, "PK": false, "AI": false, "UN": true}]}]`
	var waaaa []data
	json.Unmarshal([]byte(tableData), &waaaa)

	for _, s := range waaaa {
		finalTable := make([]string, len(s.Tables))
		tempLastRow := ""

		for _, y := range s.Tables {
			fields := reflect.TypeOf(y)
			values := reflect.ValueOf(y)
			num := fields.NumField()

			temp := [2]string{"", ""}
			tempPk := ""

			for i := 0; i < num; i++ {
				field := fields.Field(i)
				value := values.Field(i)

				if value.Kind() == reflect.String {
					switch field.Name {
					case "Name":
						temp[0] += value.String() + " "
						tempPk = value.String()
					case "Type":
						temp[0] += value.String() + " "
					}
				} else {
					if value.Bool() == true {
						switch field.Name {
						case "PrimaryKey":
							temp[1] += "PRIMARY KEY (" + tempPk
						case "NotNull":
							temp[0] += "NOT NULL "
						case "AutoIncrement":
							temp[1] += " AUTOINCREMENT)"
						case "Unique":
							temp[0] += "UNIQUE"
						}
					}
				}
			}

			temp[0] += ","
			if temp[1] != "" && temp[1][len(temp[1])-1:] != ")" {
				temp[1] += ")"
			}

			finalTable = append(finalTable, temp[0])
			tempLastRow += temp[1]
		}
		tempFinal := ""
		for _, t := range finalTable {
			tempFinal += t
		}

		query := fmt.Sprintf("CREATE TABLE %s (%s %s);", waaaa[0].TableName, tempFinal, tempLastRow)
		fmt.Println(query)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//queryTest := "UPDATE " + tableName + " SET " + field + " = '" + fieldValue + "' WHERE " + primary + " = '" + primaryValue + "'"
//fmt.Println(queryTest)
//query := fmt.Sprint("UPDATE %d SET %d = %d WHERE %d = %d", tableName, field, fieldValue, primary, primaryValue)
// _, err = db.NamedExec(`UPDATE :table SET :column = :columnvalue WHERE :pk = :pkvalue`,
// 	map[string]interface{}{
// 		"table":       tableName,
// 		"column":      field,
// 		"columnvalue": fieldValue,
// 		"pk":          primary,
// 		"pkvalue":     primaryValue,
// 	})
// if err != nil {
// 	fmt.Println(err)
// }

//test := `UPDATE ` + tableName + `SET title = :columnvalue WHERE :pk = :pkvalue`

// test2 := "UPDATE " + tableName + " SET " + field + " = :columnvalue WHERE " + primary + " = :pkvalue"
// _, err = db.NamedExec(test2,
// 	map[string]interface{}{
// 		"columnvalue": fieldValue,
// 		"pkvalue":     primaryValue,
// 	})
// if err != nil {
// 	fmt.Println(err)
// }
// type headers struct {
// 	Cid       string         `db:"cid"`
// 	Name      string         `db:"name"`
// 	Type      string         `db:"type"`
// 	NotNull   int            `db:"notnull"`
// 	DfltValue sql.NullString `db:"dflt_value"`
// 	Pk        int            `db:"pk"`
// }

// type tableInfo struct {
// 	Cid       int            `db:"cid"`
// 	Name      string         `db:"name"`
// 	Type      string         `db:"type"`
// 	NotNull   bool           `db:"notnull"`
// 	DfltValue sql.NullString `db:"dflt_value"`
// 	Pk        bool           `db:"pk"`
// }

// results := []tableInfo{}
// query := "PRAGMA table_info(albums)"

// errQuery := db.Select(&results, query)
// if errQuery != nil {
// 	fmt.Println(errQuery)
// }

// fmt.Println(results)

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
