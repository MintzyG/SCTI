package main

import (
  "SCTI/fileserver"
  "SCTI/middleware"
  "log"
  "net/http"
  "github.com/joho/godotenv"
  "github.com/lengzuo/supa"
  "os"
  "fmt"
)

func start() (*supabase.Client) {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  conf := supabase.Config{
    ApiKey:     os.Getenv("SUPABASE_KEY"), 
    ProjectRef: os.Getenv("SUPABASE_URL"),
    Debug:      false,
  }
  supaClient, err := supabase.New(conf)
  if err != nil {
      fmt.Printf("failed in initialise client with err: %s", err)
      panic("FUCK")
  }
  return supaClient
}

type SignUpRequest struct {
  Email string
  Password string
}

func main() {
  supaClient := start()
  fileserver.RunFileServer()

  mux := http.NewServeMux()
  LoadRoutes(mux, supaClient)

  server := http.Server{
    Addr:    ":8080",
    Handler: middleware.EndpointLogging(mux),
  }

  log.Fatal(server.ListenAndServe())
}
