package auth

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de cadastro e login")
}
