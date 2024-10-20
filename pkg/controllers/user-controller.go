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
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}
