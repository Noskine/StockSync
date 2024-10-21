package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Noskine/StockSync/internal/config"
	"github.com/Noskine/StockSync/internal/usecases"
	"github.com/Noskine/StockSync/pkg/dto"
)

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var data dto.InputServerUserDTO

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := config.Valid(data); err != nil {
		http.Error(w, "Validation error in the request", http.StatusBadRequest)
		return
	}

	id, err := usecases.NewUseCaseUser().CreateNewUser(data)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			http.Error(w, "User already exists or email already registered", http.StatusBadRequest)
			return
		}

		http.Error(w, "Create user Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func FindAllController(w http.ResponseWriter, r *http.Request) {
	result, err := usecases.NewUseCaseUser().GetUsers()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal server error / formater", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "JSON")
	w.Write(js)
}
