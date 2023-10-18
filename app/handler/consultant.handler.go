package handler

import (
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/app/service"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetSpecializationCategories(ctx *fiber.Ctx) error {
	return ctx.JSON(service.GetAllCategories())
}

func GetAllConsultants(ctx *fiber.Ctx) error {
	consultants, err := service.GetAllConsultants()
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(consultants)
}

func GetConsultantsBySpecializations(ctx *fiber.Ctx) error {
	specialization := ctx.Params("specialization")

	consultants, err := service.GetConsultantsBySpecialization(specialization)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(consultants)
}

func AddConsultant(ctx *fiber.Ctx) error {
	var body model.AddConsultantParam
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

	// add consultant
	user, err := service.AddConsultant(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
}

func UpdateConsultant(ctx *fiber.Ctx) error {
	var body model.UpdateConsultantParam
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

	// update consultant
	user, err := service.UpdateConsultant(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
}

func DeleteConsultant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := service.RemoveConsultant(id)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "Deleted successfully",
	})
}
