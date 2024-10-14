package repository_test

import (
	"log"
	"testing"

	"github.com/Noskine/StockSync/internal/entities"
	inmemory "github.com/Noskine/StockSync/internal/repository/in-memory"
	"github.com/google/uuid"
)

var id string

func TestInMemory(t *testing.T) {
	t.Run("testing the create user function in memory", func(t *testing.T) {
		repository := inmemory.NewInMemoryPessoaRepository()

		_, err := repository.Create(entities.User{
			Id:   uuid.NewString(),
			Name: "Enikson Sonay",
		})
		if err != nil {
			t.Error(err)
		}

		id = repository.NextID

		log.Println(repository.FindAll())
	})

	t.Run("testing the delete-all user function in memory", func(t *testing.T) {
		repository := inmemory.NewInMemoryPessoaRepository()

		err := repository.DeleteAll()
		if err != nil {
			t.Error(err)
		}

		log.Println(repository.FindAll())
	})
}
