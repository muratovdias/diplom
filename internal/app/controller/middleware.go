package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const ctxUserKey ctxKey = iota

type ctxKey int8

func (h *Handler) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil || sessionID.Value == "" {
			// fmt.Println("middleware", err.Error())
			next.ServeHTTP(w, r)
			return
		} else {
			user, err := h.service.AuthService.GetUserByToken(sessionID.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			if user.TokenDuration.Before(time.Now()) {
				fmt.Println("middleware token expired", sessionID.Expires)
				http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, user)))
		}
	})
}
