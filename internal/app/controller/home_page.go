package controller

import (
	"net/http"

	"github.com/muratovdias/diplom/internal/models"
)

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(ctxUserKey).(models.User)
	switch user.Role {
	case "trainer":
		trainings, err := h.service.Trainer.ViewAllTrainings(user.ID)
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
	default:
		trainers, err := h.service.Client.GetAllTrainers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mainPage := models.MainPage{
			Trainers: trainers,
			User:     user,
		}
		if err := h.templExecute(w, "./ui/client_home.html", mainPage); err != nil {
			return
		}
	}
}
