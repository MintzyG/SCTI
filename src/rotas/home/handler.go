package home

import (
	"SCTI/fileserver"
	"net/http"
)

type Handler struct{}

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	fileserver.T = fileserver.Execute("template/index.gohtml")
	fileserver.T.Execute(w, nil)
}
