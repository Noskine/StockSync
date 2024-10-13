package methods

import (
	"github.com/Noskine/StockSync/internal/entities"
	"github.com/Noskine/StockSync/pkg/database"
)

func InsetUser(user entities.User) (id string, err error) {
	conn := database.OpenConn()
	defer conn.Close()

	strQuery := "INSERT INTO users (id, name) VALUES ($1, $2) RETURNING id;"

	if err := conn.QueryRow(strQuery, user.Id, user.Name).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
