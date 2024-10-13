package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect(str string) {
	d, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}
	Db = d
}
