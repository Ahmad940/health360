package router

import (
	"github.com/Ahmad940/health360/app/handler"
	"github.com/gofiber/fiber/v2"
)

func USSD(app fiber.Router) {
	auth := app.Group("/ussd")

	auth.Post("/", handler.StartUSSD)

}
