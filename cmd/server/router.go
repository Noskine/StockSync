package server

import (
	"net/http"

	"github.com/Noskine/StockSync/pkg/controllers"
)

func Router() {
	r := http.NewServeMux()

	r.Handle("GET /", &controllers.UserController{})

	appStart(r)
}
