package server

import (
	"net/http"

	"github.com/Noskine/StockSync/internal/config"
)

func appStart(mux *http.ServeMux) {
	env := config.Env

	if err := http.ListenAndServe(env.Port, mux); err != nil {
		panic("erro ao iniciar o servidor HTTPS")
	}
}
