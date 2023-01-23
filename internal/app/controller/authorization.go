package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/diplom/internal/models"
)

var id int
var err error

func (h *Handler) SignInGet(ctx *fiber.Ctx) error {
	trainer := new(models.Trainer)
	if err := ctx.BodyParser(trainer); err != nil {
		fmt.Println("error = ", err)
		return ctx.SendStatus(400)
	}
	fmt.Println(trainer)
	return ctx.Status(200).SendString(trainer.Name)
}

func (h *Handler) SignInPost(ctx *fiber.Ctx) error {

	return nil
}

func (h *Handler) SignUpGet(ctx *fiber.Ctx) error {
	trainer := new(models.Trainer)
	if err := ctx.BodyParser(trainer); err != nil {
		fmt.Println("error = ", err)
		return ctx.SendStatus(400)
	}
	id, err = h.service.CreateTrainer(trainer)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
	return ctx.Status(200).SendString("success")
}

func (h *Handler) SignUpPost(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	fmt.Println(role)
	switch role {
	case "trainer":
		trainer := new(models.Trainer)
		if err = ctx.BodyParser(trainer); err != nil {
			fmt.Println("error", err)
			return ctx.SendStatus(400)
		}
		id, err = h.service.CreateTrainer(trainer)
		if err != nil {
			fmt.Println(err.Error())
			return ctx.SendStatus(500)
		}
	case "client":
		client := new(models.Client)
		if err = ctx.BodyParser(client); err != nil {
			fmt.Println("error:", err)
			return ctx.SendStatus(400)
		}
		id, err = h.service.CreateClient(client)
		if err != nil {
			fmt.Println(err.Error())
			return ctx.SendStatus(500)
		}
	default:
		return ctx.SendStatus(400)
	}

	fmt.Println(id)
	return ctx.Status(200).SendString("success")
}
