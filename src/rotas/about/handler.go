package about

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sobre")
}
