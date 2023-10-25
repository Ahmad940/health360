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

	fmt.Println("Body", string(ctx.Body()))

	res := service.StartUSSD(sessionId, serviceCode, phoneNumber, text)

	ctx.Set("Content-Type", "text/plain")
	return ctx.Send([]byte(res))
}
