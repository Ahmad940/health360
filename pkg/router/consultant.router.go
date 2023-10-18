package router

import (
	"github.com/Ahmad940/health360/app/handler"
	"github.com/gofiber/fiber/v2"
)

func Consultant(app fiber.Router) {
	auth := app.Group("/consultant")

	auth.Get("/", handler.GetAllConsultants)
	auth.Get("/categories", handler.GetSpecializationCategories)
	auth.Get("/:specialization", handler.GetConsultantsBySpecializations)
	auth.Post("/", handler.AddConsultant)
	auth.Patch("/", handler.UpdateConsultant)
	auth.Delete("/:id", handler.DeleteConsultant)
}
