package eventos

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetEventos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Atendentes, palestras e minicursos")
}
