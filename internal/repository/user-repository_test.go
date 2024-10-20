package repository_test

import (
	"fmt"
	"testing"

	"github.com/Noskine/StockSync/internal/entities"
	"github.com/Noskine/StockSync/internal/interfaces"
	inmemory "github.com/Noskine/StockSync/internal/repository/inMemory"
	"github.com/google/uuid"
)

var rep interfaces.UserRepositoryInterface

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func TestInMemory(t *testing.T) {
	rep = inmemory.NewInMemoryUserRepository()

	var id string

	t.Run("testing the create user function in memory ", func(t *testing.T) {
		str, err := rep.Create(entities.User{
			Id:   uuid.NewString(),
			Name: "Enikson Sonay R. Aires",
		})

		if err != nil {
			t.Error(err)
		}

		fmt.Println(Green + str + Reset)

		strTwo, err := rep.Create(entities.User{
			Id:   uuid.NewString(),
			Name: "Xaio Victor Duarte",
		})

		id = strTwo

		fmt.Println(Green + strTwo + Reset)
	})

	t.Run("testing the findall user function in memory", func(t *testing.T) {

		users, err := rep.FindAll()
		if err != nil {
			t.Error(err)
		}

		str := fmt.Sprintf(Yellow+"%v"+Reset, users)

		fmt.Println(str)
	})

	t.Run("testing the findbyid user function in memory", func(t *testing.T) {

		user, err := rep.FindById(id)
		if err != nil {
			t.Error(err)
		}

		str := fmt.Sprintf(Green+"%v"+Reset, user)

		fmt.Println(str)
	})

	t.Run("testing the deletebyid user function in memory", func(t *testing.T) {

		user, err := rep.DeleteById(id)
		if err != nil {
			t.Error(err)
		}

		str := fmt.Sprintf(Red+"%v"+Reset, user)

		fmt.Println(str)
	})
}
