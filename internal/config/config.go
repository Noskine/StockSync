package config

import (
	"os"

	"github.com/Noskine/StockSync/pkg/database"
	"github.com/joho/godotenv"
)

type PreInitialization struct {
	Port         string
	Host         string
	ConnectStrDB string
}

var Env = new(PreInitialization)

func init() {
	godotenv.Load()
	Env.Port = os.Getenv("PORT")
	Env.Host = os.Getenv("HOST")
	Env.ConnectStrDB = os.Getenv("StrDBConnect")

	database.Connect(Env.ConnectStrDB)
}
