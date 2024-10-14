package main

import (
	"fmt"

	_ "github.com/Noskine/StockSync/internal/config"
	"github.com/Noskine/StockSync/internal/repository"
)

func main() {
	rep := new(repository.UserRepository)

	result, err := rep.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
