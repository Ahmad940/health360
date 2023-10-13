package router

import (
	"github.com/Ahmad940/health360/app/handler"
	"github.com/Ahmad940/health360/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func User(app fiber.Router) {
	auth := app.Group("/user")

	auth.Get("/:id", handler.GetAUser)
	auth.Get("/", handler.GetAllUsers)
	auth.Patch("/", handler.UpdateUser)
	auth.Patch("/admin/update", middleware.JWTProtected(), middleware.RoleAuthorization(middleware.RoleConfig{
		Roles: []string{"admin"},
	}), handler.UpdateUserAdmin)
}
