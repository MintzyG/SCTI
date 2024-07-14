package middleware

import (
  "os"
  "log"
  "fmt"
  "net/http"
  "time"

  "github.com/golang-jwt/jwt/v5"
  "github.com/lengzuo/supa"
)

type Handler struct {
  S *supabase.Client
}

type wrappedWriter struct {
  http.ResponseWriter
  statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
  w.ResponseWriter.WriteHeader(statusCode)
  w.statusCode = statusCode
}

func EndpointLogging(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    wrapped := &wrappedWriter{
      ResponseWriter: w,
      statusCode:     http.StatusOK,
    }

    h.ServeHTTP(wrapped, r)
    log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
    })
}

func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
  if err != nil {
    fmt.Println("Cookie retrieval error:", err)
    return false
  }

  tokenString := cookie.Value
  fmt.Println("Token retrieved:", tokenString)
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return []byte(os.Getenv("SUPABASE_JWT")), nil
    })

  if err != nil {
    fmt.Println("JWT parsing error:", err)
    return false
  }

  if !token.Valid {
    fmt.Println("JWT is invalid")
    return false
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    fmt.Println("Failed to parse claims")
    return false
  }

  expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
  if expirationTime.Before(time.Now()) {
    fmt.Println("Token expired")
    return false
  }

  return true
}

