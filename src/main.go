package main

import (
	"SCTI/fileserver"
	"SCTI/middleware"
	"log"
	"net/http"
)

func main() {
	fileserver.RunFileServer()

	mux := http.NewServeMux()
	LoadRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.EndpointLogging(mux),
	}

	log.Fatal(server.ListenAndServe())
}
