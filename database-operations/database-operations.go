package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const create string = `
CREATE TABLE IF NOT EXISTS example (
id INTEGER NOT NULL PRIMARY KEY,
time DATETIME NOT NULL,
description TEXT
)`

const databaseFile string = "example.db"

func main() {
	e, _ := NewExample()
	_, insertErr := e.Insert(ExampleRow{Description: "This is a description2", Time: time.Now()})
	if insertErr != nil {
		return
	}
	all, _ := e.SelectAll()
	allBytes, _ := json.Marshal(all)
	fmt.Println(string(allBytes))
	byId, _ := e.SelectById(2)
	byIdBytes, _ := json.Marshal(byId)
	fmt.Println(string(byIdBytes))

}

type ExampleTable struct {
	db *sql.DB
}
type ExampleRow struct {
	Id          int       `json:"id"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

func (e *ExampleTable) Insert(row ExampleRow) (int, error) {
	res, err := e.db.Exec("INSERT INTO example VALUES (NULL,?,?);", row.Time, row.Description)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

type SelectOpts struct {
	key   string
	value interface{}
}

func (e *ExampleTable) SelectById(id int) (ExampleRow, error) {
	row := e.db.QueryRow("SELECT * FROM example WHERE id=?", id)

	var example ExampleRow

	err := row.Scan(&example.Id, &example.Time, &example.Description)
	return example, err
}
func (e *ExampleTable) SelectAll() ([]ExampleRow, error) {
	var err error
	var rows *sql.Rows
	rows, err = e.db.Query("SELECT * FROM example")
	// parse the rows
	var data []ExampleRow
	for rows.Next() {
		i := ExampleRow{}
		err = rows.Scan(&i.Id, &i.Time, &i.Description)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}

// NewExample Create a new table.
func NewExample() (*ExampleTable, error) {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &ExampleTable{
		db: db,
	}, nil
}
