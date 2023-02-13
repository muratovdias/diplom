package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/diplom/internal/app/service"
	"github.com/muratovdias/diplom/internal/models"
)

// type statusResponse struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// }

var id int
var err error

func (h *Handler) SignInGet(ctx *fiber.Ctx) error {
	return ctx.SendFile("ui/sign-in.html")
}

func (h *Handler) SignInPost(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	signIn := new(models.SignIn)

	if role != "trainer" && role != "client" {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "no such role",
		})
	}
	if err = ctx.BodyParser(signIn); err != nil {
		fmt.Println("error", err)
		return ctx.SendStatus(400)
	}
	err = h.service.Authentication(role, signIn.Email, signIn.Password)
	if err != nil {
		switch err {
		case service.ErrUnautorized:
			return ctx.Status(401).SendString(err.Error())
		default:
			return ctx.Status(400).SendString(err.Error())
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Success",
	})
}

func (h *Handler) SignUpGet(ctx *fiber.Ctx) error {
	// Render the HTML file
	return ctx.SendFile("ui/sign-up.html")
}

func (h *Handler) SignUpPost(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	user := new(models.User)
	switch role {
	case "trainer":
		body := ctx.Body()
		if err = json.Unmarshal(body, user); err != nil {
			return ctx.Status(400).Send([]byte(err.Error()))
		}
		id, err = h.service.CreateTrainer(user)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status":  500,
				"message": err.Error(),
			})
		}
	case "client":
		body := ctx.Body()
		if err = json.Unmarshal(body, user); err != nil {
			return ctx.Status(500).SendString(err.Error())
		}
		id, err = h.service.CreateClient(user)
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}
	default:
		return ctx.Status(400).SendString("No such role")
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"id":     id,
	})
}

func (h *Handler) Home(ctx *fiber.Ctx) error {
	return ctx.SendFile("ui/home.html")
}
