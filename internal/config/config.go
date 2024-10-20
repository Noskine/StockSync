package config

import (
	"os"

	"github.com/Noskine/StockSync/pkg/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PreInitialization struct {
	Port string
	Host string
}

var Env = new(PreInitialization)

func init() {
	godotenv.Load()
	Env.Port = os.Getenv("PORT")
	Env.Host = os.Getenv("HOST")

	database.Migrate()
}


