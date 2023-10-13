package handler

import (
	"fmt"

	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/app/service"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func RequestOTP(ctx *fiber.Ctx) error {
	var body model.Auth
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Requesting OTP
	err = service.RequestOTP(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}
	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("OTP has been sent to %v%v successfully", body.CountryCode, body.PhoneNumber),
	})
}

func Login(ctx *fiber.Ctx) error {
	var body model.Login

	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body
	response, err := service.Login(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(response)
}
