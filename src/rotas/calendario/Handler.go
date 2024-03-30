package calendario

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetCalendario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Calendário")
}
