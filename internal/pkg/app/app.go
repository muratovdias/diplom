package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	router     *chi.Mux
	config     *config.DBConfig
}

func NewApp() *App {
	var app App
	app.config, err = config.GetConfig()
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	db, err := repository.IniitDB(app.config)
	if err != nil {
		log.Fatal(err)
	}
	// if err = repository.CreateTables(db); err != nil {
	// 	log.Fatal(err)
	// }
	tx := db.MustBegin()
	app.repository = repository.NewRepository(db, tx)
	app.service = service.NewService(app.repository)
	app.handler = controller.NewHandler(app.service)

	return &app
}

func (app *App) Run() error {
	app.router = controller.InitRoutes(app.handler) //Initialize routes
	err := http.ListenAndServe(":8888", app.router)
	if err != nil {
		return err
	}
	return nil
}
