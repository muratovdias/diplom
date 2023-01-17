package main

import (
	"log"

	"github.com/muratovdias/diplom/internal/pkg/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
