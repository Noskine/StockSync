package usecases

import (
	"github.com/Noskine/StockSync/internal/entities"
	"github.com/Noskine/StockSync/internal/repository"
	"github.com/Noskine/StockSync/pkg/dto"
	"github.com/Noskine/StockSync/pkg/lib"
	"github.com/google/uuid"
)

type UseCaseUser struct {
	// Usando o repositorio de comunicação com o DB aqui. Para chamar nos metodos da struct;
	rep repository.UserRepository
}

func NewUseCaseUser() *UseCaseUser {
	return &UseCaseUser{}
}

func (u *UseCaseUser) CreateNewUser(dto dto.InputServerUserDTO) (string, error) {

	str, err := u.rep.Create(entities.User{
		Id:       uuid.NewString(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: lib.HashingPasswordFunc(dto.Pass),
	})

	if err != nil {
		return "", err
	}

	return str, nil
}

func (u *UseCaseUser) GetUsers() ([]dto.OutPutServerUserDTO, error) {
	users, err := u.rep.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []dto.OutPutServerUserDTO

	for _, user := range users {
		var dto dto.OutPutServerUserDTO
		dto.Email = user.Email
		dto.Id = user.Id
		dto.Name = user.Name

		dtos = append(dtos, dto)
	}

	return dtos, err
}

func (u *UseCaseUser) GetUserById(id string) (dto.OutPutServerUserDTO, error) {
	user, err := u.rep.FindById(id)
	if err != nil {
		return dto.OutPutServerUserDTO{}, err
	}
	return dto.OutPutServerUserDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
