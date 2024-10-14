package inmemory

import (
	"errors"

	"github.com/Noskine/StockSync/internal/entities"
	"github.com/google/uuid"
)

type (
	InMemoryUserRepository struct {
		User   map[string]entities.User
		NextID string
	}
)

func NewInMemoryPessoaRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		User:   make(map[string]entities.User),
		NextID: uuid.NewString(),
	}
}

func (rep *InMemoryUserRepository) Create(e entities.User) (string, error) {
	e.Id = rep.NextID
	rep.User[rep.NextID] = e

	return rep.NextID, nil
}

func (rep *InMemoryUserRepository) FindAll() ([]entities.User, error) {
	var pessoas []entities.User
	for _, pessoa := range rep.User {
		pessoas = append(pessoas, pessoa)
	}
	return pessoas, nil
}

func (repo *InMemoryUserRepository) FindById(id string) (*entities.User, error) {
	if user, ok := repo.User[id]; ok {
		return &user, nil
	}

	return nil, errors.New("user not found")
}

func (rep *InMemoryUserRepository) DeleteAll() error {
	rep.User = make(map[string]entities.User)

	return nil
}
