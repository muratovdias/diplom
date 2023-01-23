package controller

import (
	"github.com/gofiber/fiber/v2"
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

func InitRoutes(h Handler) *fiber.App {
	app := fiber.New()
	//Authorization routes
	auth := app.Group("/auth")
	signIn := auth.Group("/sign-in")
	signIn.Get("/", h.SignInGet)
	signIn.Post("/:role", h.SignInPost)
	signUp := auth.Group("/sign-up")
	signUp.Get("/", h.SignUpGet)
	signUp.Post("/:role", h.SignUpPost)

	return app
}
