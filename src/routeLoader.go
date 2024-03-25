package main

import (
	"SCTI/rotas/about"
	"SCTI/rotas/auth"
	"SCTI/rotas/calendario"
	"SCTI/rotas/home"
	"SCTI/rotas/horario"
	"SCTI/rotas/loja"
	eventos "SCTI/rotas/participantes_e_eventos"
	"net/http"
)

func LoadRoutes(mux *http.ServeMux) {
	loginHandler := &auth.Handler{}
	aboutHandler := &about.Handler{}
	homeHandler := &home.Handler{}
	eventosHandler := &eventos.Handler{}
	calendarioHandler := &calendario.Handler{}
	lojaHandler := &loja.Handler{}
	horarioHandler := &horario.Handler{}

	mux.HandleFunc("GET /login", loginHandler.GetLogin)
	// mux.HandleFunc("POST /login", loginHandler.PostLogin)
	mux.HandleFunc("GET /about", aboutHandler.GetAbout)

	mux.HandleFunc("GET /", homeHandler.GetHome)

	mux.HandleFunc("GET /eventos", eventosHandler.GetEventos)

	mux.HandleFunc("GET /calendario", calendarioHandler.GetCalendario)

	mux.HandleFunc("GET /loja", lojaHandler.GetLoja)

	mux.HandleFunc("GET /horario", horarioHandler.GetHorario)
}
