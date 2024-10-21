package server

import (
	"fmt"
	"net/http"

	"github.com/Noskine/StockSync/pkg/controllers"
)

func Router() {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Olá mundo!")
	})

	r.HandleFunc("POST /user", controllers.CreateUserController)
	r.HandleFunc("GET /users", controllers.GetUsersController)
	r.HandleFunc("GET /user", controllers.GetUserController)
	appStart(r)
}
