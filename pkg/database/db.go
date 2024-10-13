package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect(str string) {
	d, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}
	Db = d

	err = RunMigrations(migrations)
	if err != nil {
		panic("Error loading migrations to database")
	}

	fmt.Println("Banco de dados migrado!")
}
