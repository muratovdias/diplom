package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/diplom/internal/app/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func InitRoutes(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui/"))))

	r.Group(func(r chi.Router) {
		r.Use(h.SessionMiddleware)
		r.Get("/", h.GetHome)
		r.Get("/trainer/set-schedule", h.GetTrainerSchedule)
		r.Post("/trainer/set-schedule", h.SetTrainerSchedule)
		r.Get("/trainer/profile/{id}", h.TrainerProfile)
		r.Get("/trainer/schedule/{id}", h.TrainerProfileSchedule)
		r.Post("/trainer/schedule/{id}", h.SetClientTraining)
		r.Get("/trainer/edit-profile", h.GetEditProfile)
		r.Post("/trainer/edit-profile", h.PostEditProfile)

		r.Get("/client/trainings", h.ClientTrainings)
		r.Post("/cancel-training/", h.CanlcelTraining)

	})
	r.Group(func(r chi.Router) {
		r.Get("/auth/sign-up", h.SignUpGet)
		r.Post("/auth/sign-up", h.SignUpPost)

		r.Get("/auth/sign-in", h.SignInGet)
		r.Post("/auth/sign-in", h.SignInPost)

		r.Get("/log-out", h.LogOut)
	})
	return r
}

func (h *Handler) templExecute(w http.ResponseWriter, path string, data interface{}) error {
	templ, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("ParseFiles()")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	if err = templ.Execute(w, data); err != nil {
		fmt.Println("templExecute()")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
