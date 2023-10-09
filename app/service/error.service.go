package service

import (
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/pkg/constant"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(err error, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
		Message: err.Error(),
	})
}

func SqlErrorNotFound(err error) bool {
	return err.Error() == constant.SqlNotFoundText
}

func SqlErrorIgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == constant.SqlNotFoundText {
		return nil
	}
	return err
}
