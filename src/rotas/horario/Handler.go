package horario

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetHorario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Monte seu horário")
}
