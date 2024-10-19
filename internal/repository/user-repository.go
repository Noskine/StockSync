package repository

import (
	"database/sql"
	"errors"
	"fmt"

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

func (ur *UserRepository) FindAll() ([]entities.User, error) {
	conn := database.OpenConn()
	defer conn.Close()

	strQuery := "SELECT * FROM public.users;"

	rows, err := conn.Query(strQuery)
	if err != nil {
		return nil, errors.New("error fetching data from the database")
	}

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) FindById(id string) (entities.User, error) {
	conn := database.OpenConn()
	defer conn.Close()

	userResponse := new(entities.User)

	strQuery := "SELECT id, name FROM public.users WHERE id = $1;"

	err := conn.QueryRow(strQuery, id).Scan(&userResponse.Id, &userResponse.Name)
	if err != nil {
		return *userResponse, err
	}

	return *userResponse, nil
}

func (ur *UserRepository) DeleteById(id string) error {
	conn := database.OpenConn()
	defer conn.Close()

	strQuery := "DELETE FROM public.users WHERE id = $1;"

	err := conn.QueryRow(strQuery, id)
	if err != nil {
		if err.Err() == sql.ErrNoRows {
			return fmt.Errorf("nenhuma pessoa encontrada com o ID %s", id)
		}
		return err.Err()
	}

	return nil
}
