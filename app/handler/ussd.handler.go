package handler

import (
	"github.com/Ahmad940/health360/app/service"
	"github.com/gofiber/fiber/v2"
)

func StartUSSD(ctx *fiber.Ctx) error {
	sessionId := ctx.Params("sessionId")
	serviceCode := ctx.Params("serviceCode")
	phoneNumber := ctx.Params("phoneNumber")
	text := ctx.Params("text")
	res := service.StartUSSD(sessionId, serviceCode, phoneNumber, text)

	return ctx.Send([]byte(res))
}
