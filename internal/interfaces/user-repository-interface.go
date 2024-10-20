package interfaces

import "github.com/Noskine/StockSync/internal/entities"

type (
	UserRepositoryInterface interface {
		Create(e entities.User) (string, error)
		FindAll() ([]entities.User, error)
		FindById(id string) (entities.User, error)
		DeleteById(id string) ([]entities.User, error)
	}
)
