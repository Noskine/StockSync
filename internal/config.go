package internal

import (
	"database/sql"
	"os"
)

type PreInitialization struct {
	Port         string
	Host         string
	ConnectStrDB string
}

var Con = new(PreInitialization)

func init() {
	Con.Port = os.Getenv("PORT")
	Con.Host = os.Getenv("HOST")
}

func Connect() *sql.DB {
	db, err := connectDataBase()
	if err != nil {
		panic(err)
	}

	return db
}

func connectDataBase() (*sql.DB, error) {
	connStr := Con.ConnectStrDB
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
