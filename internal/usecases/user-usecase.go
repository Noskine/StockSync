package usecases

import (
	"github.com/Noskine/StockSync/internal/entities"
	"github.com/Noskine/StockSync/internal/repository"
	"github.com/Noskine/StockSync/pkg/dto"
	"github.com/Noskine/StockSync/pkg/lib"
	"github.com/google/uuid"
)

type UseCaseUser struct {
}

func NewUseCaseUser() *UseCaseUser {
	return &UseCaseUser{}
}

func (u *UseCaseUser) CreateNewUser(dto dto.InputServerUserDTO) (string, error) {
	rep := new(repository.UserRepository)

	str, err := rep.Create(entities.User{
		Id:    uuid.NewString(),
		Name:  dto.Name,
		Email: dto.Email,
		Password: lib.HashingPasswordFunc(dto.Pass),
	})

	if err != nil {
		return "", err
	}

	return str, nil
}
