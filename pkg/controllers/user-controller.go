package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Noskine/StockSync/internal/config"
	"github.com/Noskine/StockSync/internal/usecases"
	"github.com/Noskine/StockSync/pkg/dto"
)

// Rotas de comunicação com o usuário;
func RouterUser(r *http.ServeMux) {
	r.HandleFunc("POST /user", CreateUserController)
	r.HandleFunc("GET /users", GetUsersController)
	r.HandleFunc("GET /user", GetUserController)
	r.HandleFunc("DELETE /user", DeleteUser)
}

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
		// Usuário já existe;
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			http.Error(w, "User already exists or email already registered", http.StatusBadRequest)
			return
		}

		http.Error(w, "Create user Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "Json")
	json.NewEncoder(w).Encode(id)
}

func GetUsersController(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Add("Content-Type", "Json")
	w.Write(js)
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	QueryId := r.URL.Query().Get("id")

	result, err := usecases.NewUseCaseUser().GetUserById(QueryId)
	if err != nil {
		// Usuário não existe no DB;
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "User is not registered in the database", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(result)
	if err != nil {

		http.Error(w, "Internal server error / formater", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "Json")
	w.Write(js)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	QueryId := r.URL.Query().Get("id")

	err := usecases.NewUseCaseUser().DeleteUserById(QueryId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "Json")
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		http.Error(w, "Internal server error / formater", http.StatusInternalServerError)
		return
	}
}
