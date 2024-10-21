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

func Migrate() {
	var migrations []Migration = []Migration{{
		Up: "CREATE TABLE IF NOT EXISTS users (id uuid NOT NULL, name varchar(100) NOT NULL, email varchar(255) UNIQUE NOT NULL, password varchar(100) NOT NULL, PRIMARY KEY(id));",
	}}

	RunMigrations(migrations)
}
