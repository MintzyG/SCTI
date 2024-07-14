package dashboard

import (
  "SCTI/fileserver"
  "SCTI/middleware"
  "net/http"

  // supabase "github.com/lengzuo/supa"
)

type Handler struct{}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
  if !middleware.IsAuthenticated(r) {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  }

  var t = fileserver.Execute("template/dashboard.gohtml")
  t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux) {
  handler := &Handler{}
  mux.HandleFunc("GET /dashboard",  handler.GetDashboard)
}
