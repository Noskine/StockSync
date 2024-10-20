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

	appStart(r)
}
