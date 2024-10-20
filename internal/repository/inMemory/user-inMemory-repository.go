package inmemory

import (
	"errors"

	"github.com/Noskine/StockSync/internal/entities"
)

type InMemoryUserRepository struct {
	users []entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

func (iM *InMemoryUserRepository) Create(e entities.User) (string, error) {

	iM.users = append(iM.users, e)

	for _, user := range iM.users {
		if user.Id == e.Id {
			return user.Id, nil
		}
	}

	return "", nil
}

func (iM *InMemoryUserRepository) FindAll() ([]entities.User, error) {
	return iM.users, nil
}

func (iM *InMemoryUserRepository) FindById(id string) (entities.User, error) {
	for _, user := range iM.users {
		if user.Id == id {
			return user, nil
		}
	}

	return entities.User{}, errors.New("")
}

func (iM *InMemoryUserRepository) DeleteById(id string) ([]entities.User, error) {
	for i, users := range iM.users {
		if users.Id == id {
			// Remove a users.Id, excluindo o item `i` do slice
			iM.users = append(iM.users[:i], iM.users[i+1:]...)

			return iM.users, nil
		}
	}

	return nil, errors.New("")
}
