package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func OpenConn() *sql.DB {
	d, err := sql.Open("postgres", os.Getenv("StrDBConnect"))
	if err != nil {
		panic(err)
	}
	return d
}
