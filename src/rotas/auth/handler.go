package auth

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "SCTI/fileserver"
  "github.com/lengzuo/supa"
  "github.com/lengzuo/supa/dto"
)

type Handler struct{
  S *supabase.Client
}

type User struct {
  Email string
  Password string 
}

type SupaError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
  println("In PostSignup")
  ctx := r.Context()

  var user User

  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }

  fmt.Println(user.Email)
  fmt.Println(user.Password)

  body := dto.SignUpRequest{
    Email:    user.Email,
    Password: user.Password,
  }

  _, err := h.S.Auth.SignUp(ctx, body)
  if err != nil {
    // Check if the error is of type SupaError
    var supaErr SupaError
    if jsonErr := json.Unmarshal([]byte(err.Error()), &supaErr); jsonErr == nil {
      switch supaErr.ErrorCode {
      case "user_already_exists":
        http.Error(w, "User already registered", http.StatusConflict)
        return
      case "anonymous_provider_disabled":
        http.Error(w, "Anonymous sign-ins are disabled", http.StatusForbidden)
        return
      default:
        http.Error(w, fmt.Sprintf("Error signing up: %v", supaErr.Msg), http.StatusInternalServerError)
        return
      }
    }
    // Handle generic error if unable to parse SupaError
    http.Error(w, fmt.Sprintf("Request error: %v", err), http.StatusInternalServerError)
    return
  }

  // Read the successful response if needed
  // var authDetail AuthDetailResp
  // bodyBytes, readErr := ioutil.ReadAll(resp.Body)
  // if readErr != nil {
  //     http.Error(w, fmt.Sprintf("Failed to read response body: %v", readErr), http.StatusInternalServerError)
  //     return
  // }
  // jsonErr := json.Unmarshal(bodyBytes, &authDetail)
  // if jsonErr != nil {
  //     http.Error(w, fmt.Sprintf("Failed to parse success response: %v", jsonErr), http.StatusInternalServerError)
  //     return
  // }

  // Redirect to the lncc page after successful signup
  http.Redirect(w, r, "/lncc", http.StatusSeeOther)
}

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  println("In GetSignup")
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Pagina de login")
}

func RegisterRoutes(mux *http.ServeMux, s *supabase.Client) {
  handler := &Handler{S: s}
  mux.HandleFunc("GET /signup", handler.GetSignup)
  mux.HandleFunc("POST /signup", handler.PostSignup)
  mux.HandleFunc("/login", handler.GetLogin)
}
