package main

import (
	"github.com/Noskine/StockSync/cmd/server"
	_ "github.com/Noskine/StockSync/internal/config"
)

func main() {
	server.Router();
}
