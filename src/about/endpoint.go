package about

import (
  "net/http"
  "fmt"
)

func Endpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%v: Hello", r.Header) 
}
