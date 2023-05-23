package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/diplom/internal/models"
)

func (h *Handler) SetClientTraining(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	idStr := chi.URLParam(r, "id")
	trainerID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var times []string
	r.ParseForm()
	for i := 0; i < 7; i++ {
		date := time.Now().Add(time.Hour * time.Duration(24*i)).Format("2006-01-02")
		time, ok := r.Form[date]
		if ok {
			for _, t := range time {
				d := fmt.Sprintf("%s %s", date, t)
				times = append(times, d)
			}
		}
	}
	note := r.FormValue("note")
	if len(note) == 0 || strings.TrimSpace(note) == "" {
		note = "NO"
	}
	if err := h.service.Client.SetTrainings(user.ID, trainerID, note, times); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/trainer/schedule/"+idStr, http.StatusSeeOther)
}

func (h *Handler) ClientTrainings(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	trainings, err := h.service.Client.ViewTrainings(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageDate := models.TrainerHomePage{
		User:      user,
		Trainings: trainings,
	}
	if err := h.templExecute(w, "./ui/trainings.html", pageDate); err != nil {
		return
	}
}

func (h *Handler) CanlcelTraining(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	date := r.URL.Query().Get("date")
	// fmt.Println(user.ID, date)
	if err := h.service.Client.CancelTraining(user.ID, date); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/client/trainings", http.StatusSeeOther)
}

func (h *Handler) MyStats(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	stats, err := h.service.Client.MyStats(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stats.User = user
	if err := h.templExecute(w, "./ui/my-stats.html", stats); err != nil {
		return
	}
}

func (h *Handler) LikeTrainer(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	id := chi.URLParam(r, "id")
	trainerId, err := strconv.Atoi(id)
	// fmt.Println(trainerId)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if err := h.service.Client.LikeTrainer(user.ID, trainerId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) DislikeTrainer(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	id := chi.URLParam(r, "id")
	trainerId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if err := h.service.Client.DislikeTrainer(user.ID, trainerId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
