package server

import "net/http"

func Router() {
	r := http.NewServeMux()

	r.HandleFunc("/ola", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo!"))
	})
	
	appStart(r)
}
