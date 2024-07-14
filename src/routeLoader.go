package main

import (
	"SCTI/fileserver"
	"SCTI/rotas/about"
	"SCTI/rotas/auth"
	"SCTI/rotas/calendario"
	"SCTI/rotas/home"
	"SCTI/rotas/horario"
	"SCTI/rotas/lncc"
	"SCTI/rotas/loja"
  "SCTI/rotas/dashboard"
	eventos "SCTI/rotas/participantes_e_eventos"
	"net/http"

	supabase "github.com/lengzuo/supa"
)

func LoadRoutes(mux *http.ServeMux, s *supabase.Client) {
  mux.Handle("/static/", http.StripPrefix("/static/", fileserver.FS))

	aboutHandler := &about.Handler{}
	eventosHandler := &eventos.Handler{}
	calendarioHandler := &calendario.Handler{}
	lojaHandler := &loja.Handler{}
	horarioHandler := &horario.Handler{}

  auth.RegisterRoutes(mux, s)
  dashboard.RegisterRoutes(mux)
  home.RegisterRoutes(mux)
  lncc.RegisterRoutes(mux)

  mux.HandleFunc("GET /about", aboutHandler.GetAbout)
	mux.HandleFunc("GET /eventos", eventosHandler.GetEventos)
	mux.HandleFunc("GET /calendario", calendarioHandler.GetCalendario)
	mux.HandleFunc("GET /loja", lojaHandler.GetLoja)
	mux.HandleFunc("GET /horario", horarioHandler.GetHorario)
}
