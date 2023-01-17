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

func InitRoutes() *fiber.App {
	app := fiber.New()
	//Auth routes
	auth := app.Group("/auth")
	signIn := auth.Group("/sign-in")
	signIn.Get("/", SignIn)
	// signIn.Post("/client")
	// signIn.Post("/trainer")
	signUp := auth.Group("/sign-up")
	signUp.Get("/", SignUp)
	// signIn.Post("/client")
	// signIn.Post("/trainer")

	return app
}

func SignIn(ctx *fiber.Ctx) error {
	// var trainer models.Trainer

	return ctx.Status(200).SendString("Sign In page")
}

func SignUp(ctx *fiber.Ctx) error {
	// ctx.Status(200).SendString("Sign Up page")
	return ctx.Status(200).SendString("Sign Up page")
}
