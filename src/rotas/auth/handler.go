package auth

import (
	"SCTI/fileserver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "time"
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
  Error     string `json:"error"`
  ErrorDesc string `json:"error_description"`
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
        if supaErr.Msg != "" {
          http.Error(w, fmt.Sprintf("Error signing up: %s", supaErr.Msg), http.StatusInternalServerError)
        } else {
          http.Error(w, fmt.Sprintf("%s: %s", supaErr.Error, supaErr.ErrorDesc), http.StatusInternalServerError)
        }
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
  http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  println("In GetSignup")
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
  fmt.Println("In PostLogin")
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

  body := dto.SignInRequest{
    Email:    user.Email,
    Password: user.Password,
  }

  resp, err := h.S.Auth.SignInWithPassword(ctx, body)
  if err != nil {
    // Check if the error is of type SupaError
    var supaErr SupaError
    if jsonErr := json.Unmarshal([]byte(err.Error()), &supaErr); jsonErr == nil {
      switch supaErr.ErrorCode {
      case "invalid_credentials":
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
      case "user_not_found":
        http.Error(w, "User not found", http.StatusNotFound)
        return
      case "anonymous_provider_disabled":
        http.Error(w, "Anonymous sign-ins are disabled", http.StatusForbidden)
        return
      default:
        if supaErr.Msg != "" {
          http.Error(w, fmt.Sprintf("Error signing up: %s", supaErr.Msg), http.StatusInternalServerError)
        } else {
          http.Error(w, fmt.Sprintf("%s: %s", supaErr.Error, supaErr.ErrorDesc), http.StatusInternalServerError)
        }
        return
      }
    }
    // Handle generic error if unable to parse SupaError
    http.Error(w, fmt.Sprintf("Request error: %v", err), http.StatusInternalServerError)
    return
  }

  authCookie := &http.Cookie{
      Name:     "auth",
      Value:    resp.AccessToken,
      Path:     "/",
      HttpOnly: true,
      Secure:   true,
      Expires:  time.Now().Add(3 * 24 * time.Hour), // Adjust the expiration time as needed
  }

  fmt.Println(resp.AccessToken)
  http.SetCookie(w, authCookie)

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
  http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  println("In GetLogin")
  var t = fileserver.Execute("template/login.gohtml")
  t.Execute(w, nil)
}

func RegisterRoutes(mux *http.ServeMux, s *supabase.Client) {
  handler := &Handler{S: s}

  mux.HandleFunc("GET /signup", handler.GetSignup)
  mux.HandleFunc("GET /login", handler.GetLogin)

  mux.HandleFunc("POST /signup", handler.PostSignup)
  mux.HandleFunc("POST /login", handler.PostLogin)
}
