package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/diplom/internal/app/service"
	"github.com/muratovdias/diplom/internal/models"
)

type Schedule struct {
	Dates map[string][]string `json:"dates"`
}

func (h *Handler) GetTrainerSchedule(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	switch user.Role {
	case "trainer":
		if err := h.templExecute(w, "./ui/trainer_set-schedule.html", user); err != nil {
			return
		}
	default:
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
}

func (h *Handler) SetTrainerSchedule(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(ctxUserKey).(models.User)
	if !ok || user.Role != "trainer" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	var schedule Schedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println("schedule", schedule)
	if err := h.service.Trainer.SetSchedule(user.ID, schedule.Dates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp models.Response
	resp.Code = http.StatusCreated
	resp.Message = "schedule set seccessfully"
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) TrainerProfile(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	idStr := chi.URLParam(r, "id")
	// fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	trainer, err := h.service.Trainer.GetFullTrainerInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := h.service.Commet.GetCommentsByTrainerID(user.ID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(comments[0].Edit)
	trainerProfile := models.TrainerProfile{
		User:        user,
		TrainerInfo: trainer,
		Comments:    comments,
	}
	if err := h.templExecute(w, "./ui/trainer_profile.html", trainerProfile); err != nil {
		return
	}
}

func (h *Handler) TrainerProfileSchedule(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	if user.Role == "" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	schedule, err := h.service.Client.ViewSchedule(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(schedule)
	pageData := models.TrainerSchedule{
		TrainerID: id,
		User:      user,
		Schedule:  schedule,
	}
	if err := h.templExecute(w, "./ui/trainer_schedule.html", pageData); err != nil {
		return
	}
}

func (h *Handler) GetEditProfile(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	if user.Role != "trainer" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	if err := h.templExecute(w, "./ui/trainer_edit_profile.html", nil); err != nil {
		return
	}
}

func (h *Handler) PostEditProfile(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var trainer models.TrainerInfo
	trainer.ID = user.ID
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	trainer.FullName = firstName + " " + lastName
	trainer.Phone = r.FormValue("phone")
	trainer.Adress = r.FormValue("adress")
	trainer.Speciality = r.FormValue("speciality")
	trainer.Bio = r.FormValue("bio")
	_, handler, _ := r.FormFile("ava")

	img, err := service.CreateImage(handler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	trainer.Img = template.URL(img)
	trainer.Facebook = r.FormValue("fb")
	trainer.Twitter = r.FormValue("tw")
	trainer.Instagram = r.FormValue("insta")
	if err := h.service.Trainer.UpdateTrainerInfo(trainer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/trainer/profile/"+strconv.Itoa(user.ID), http.StatusSeeOther)
}

func (h *Handler) CancelSchedule(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	if user.Role != "trainer" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
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
	if err := h.service.Trainer.CancelSchedule(user.ID, times); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/trainer/schedule/"+strconv.Itoa(user.ID), http.StatusSeeOther)
}
