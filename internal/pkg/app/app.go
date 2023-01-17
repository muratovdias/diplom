package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/diplom/internal/app/controller"
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/app/service"
)

type App struct {
	repository *repository.Repository
	service    *service.Service
	handler    *controller.Handler
	fiber      *fiber.App
}

func NewApp() *App {
	var app App
	app.repository = repository.NewRepository()
	app.service = service.NewService(app.repository)
	app.handler = controller.NewHandler(app.service)

	return &app
}

func (app *App) Run() error {
	app.fiber = controller.InitRoutes() //Initialize routes
	err := app.fiber.Listen(":8888")
	if err != nil {
		return err
	}
	return nil
}
