package repository

import (
	"errors"

	"github.com/Noskine/StockSync/internal/entities"
	"github.com/Noskine/StockSync/pkg/database"
)

type (
	UserRepository struct {
	}
)

func (ur *UserRepository) Create(e entities.User) (string, error) {
	conn := database.OpenConn()
	defer conn.Close()

	var id string

	strQuery := "INSERT INTO users (id, name) VALUES ($1, $2) RETURNING id;"

	if err := conn.QueryRow(strQuery, e.Id, e.Name).Scan(&id); err != nil {
		return "", errors.New("error creating user in database")

	}

	return id, nil
}
