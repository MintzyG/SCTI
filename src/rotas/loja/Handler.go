package loja

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetLoja(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Loja")
}
