package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/diplom/internal/app/config"
	"github.com/muratovdias/diplom/internal/app/controller"
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/app/service"
)

var err error

type App struct {
	repository *repository.Repository
	service    *service.Service
	handler    *controller.Handler
	fiber      *fiber.App
	config     *config.GlobalConfig
}

func NewApp() *App {
	var app App
	app.config, err = config.GetConfig()
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	if err = repository.IniitDB(app.config); err != nil {
		log.Fatal(err)
	}
	app.repository = repository.NewRepository()
	app.service = service.NewService(app.repository)
	app.handler = controller.NewHandler(app.service)

	return &app
}

func (app *App) Run() error {
	app.fiber = controller.InitRoutes(*app.handler) //Initialize routes
	err := app.fiber.Listen(app.config.Port)
	if err != nil {
		return err
	}
	return nil
}
