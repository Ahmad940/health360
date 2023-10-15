package router

import (
	"github.com/Ahmad940/health360/app/handler"
	"github.com/gofiber/fiber/v2"
)

func Authentication(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Get("/profile", handler.Profile)
	auth.Post("/request-otp", handler.RequestOTP)
	auth.Post("/login", handler.Login)
}
