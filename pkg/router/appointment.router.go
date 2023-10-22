package router

import (
	"github.com/Ahmad940/health360/app/handler"
	"github.com/Ahmad940/health360/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Appointment(app fiber.Router) {
	appointment := app.Group("/appointment")

	appointment.Get("/", middleware.JWTProtected(), handler.GetUserAppointments)
	appointment.Get("/:id", middleware.JWTProtected(), handler.GetAppointmentById)
	appointment.Post("/", middleware.JWTProtected(), handler.CreateAppointment)
	appointment.Patch("/:id", middleware.JWTProtected(), handler.UpdateAppointment)
	appointment.Delete("/:id", middleware.JWTProtected(), handler.DeleteAppointment)
}
