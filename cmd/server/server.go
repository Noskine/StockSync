package server

import (
	"net/http"
)

func appStart(mux *http.ServeMux) {

	if err := http.ListenAndServe(":3031", mux); err != nil {
		panic("erro ao iniciar o servidor HTTPS")
	}
}
