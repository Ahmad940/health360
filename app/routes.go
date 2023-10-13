package app

import (
	"github.com/Ahmad940/health360/app/handler/ws"
	"github.com/Ahmad940/health360/pkg/middleware"
	"github.com/Ahmad940/health360/pkg/router"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var cfg = websocket.Config{
	RecoverHandler: func(conn *websocket.Conn) {
		if err := recover(); err != nil {
			conn.WriteJSON(fiber.Map{"customError": "error occurred"})
		}
	},
}

func AttachRoutes(app *fiber.App) {

	base := app.Group("/")

	app.Get("/ws/:id", websocket.New(ws.Stream, cfg))

	api := base.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error { return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"}) })

	// routes
	router.Authentication(api)
	router.User(api)

	// not found
	middleware.NotFoundMiddleware(app)
}
