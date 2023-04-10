package controller

import (
	"encoding/json"
	"net/http"

	"github.com/muratovdias/diplom/internal/app/service"
	"github.com/muratovdias/diplom/internal/models"
)

// type statusResponse struct {
// 	Role string `json:"role"`
// }

func (h *Handler) SignInGet(w http.ResponseWriter, r *http.Request) {
	if err := h.templExecute(w, "./ui/sign-in.html", nil); err != nil {
		return
	}
}

func (h *Handler) SignInPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := h.service.Authentication(email, password)
	if err != nil {
		switch err {
		case service.ErrUnautorized:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	// Create new session cookie and set its value to the session token and expiration time
	sessionCookie := http.Cookie{
		Name:  "session_id",
		Value: user.Token,
		Path:  "/",
	}

	http.SetCookie(w, &sessionCookie)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (h *Handler) SignUpGet(w http.ResponseWriter, r *http.Request) {
	if err := h.templExecute(w, "./ui/sign-up.html", nil); err != nil {
		return
	}
}

func (h *Handler) SignUpPost(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := h.service.CreateUser(user)
	if err != nil {
		w.Write([]byte("status: fail"))
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: "",
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
