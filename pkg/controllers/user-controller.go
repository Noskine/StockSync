package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct{}

func (h *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from mycontroller")
}
