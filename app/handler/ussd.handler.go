package handler

import (
	"fmt"

	"github.com/Ahmad940/health360/app/service"
	"github.com/gofiber/fiber/v2"
)

func StartUSSD(ctx *fiber.Ctx) error {
	sessionId := ctx.FormValue("sessionId")
	serviceCode := ctx.FormValue("serviceCode")
	phoneNumber := ctx.FormValue("phoneNumber")
	text := ctx.FormValue("text")

	fmt.Println("Session Param 1", ctx.Params("sessionId"), ctx.Params("text"))
	fmt.Println("Session Query 2", ctx.Query("sessionId"), ctx.Query("text"))
	fmt.Println("Session GET 3", ctx.Get("sessionId"), ctx.Get("text"))

	res := service.StartUSSD(sessionId, serviceCode, phoneNumber, text)

	ctx.Set("Content-Type", "text/plain")
	return ctx.Send([]byte(res))
}
