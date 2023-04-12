package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/muratovdias/diplom/internal/models"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	if user.Role == "" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	trainerID := r.URL.Query().Get("trainerID")
	id, err := strconv.Atoi(trainerID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	commentText := r.FormValue("comment")
	if strings.TrimSpace(commentText) == "" {
		http.Redirect(w, r, "/trainer/profile/"+trainerID, http.StatusSeeOther)
		return
	}
	commet := models.Comment{
		TrainerID: id,
		AuthorID:  user.ID,
		Date:      time.Now().Format("01-02-2006 15:04:05"),
		Text:      commentText,
	}
	if err := h.service.Commet.Create(commet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/trainer/profile/"+trainerID, http.StatusSeeOther)
}
