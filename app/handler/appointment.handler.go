package handler

import (
	"github.com/Ahmad940/health360/app/model"
	"github.com/Ahmad940/health360/app/service"
	"github.com/Ahmad940/health360/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetUserAppointments(ctx *fiber.Ctx) error {
	userID, err := retrieveUserID(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	appointments, err := service.GetUserAppointments(userID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(appointments)
}

func GetAppointmentById(ctx *fiber.Ctx) error {
	userID, err := retrieveUserID(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	appointment, err := service.GetAppointmentById(ctx.Params("id"), userID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(appointment)
}

func CreateAppointment(ctx *fiber.Ctx) error {
	var body model.Appointment

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
	response, err := service.CreateAppointment(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(response)
}

func UpdateAppointment(ctx *fiber.Ctx) error {
	userID, err := retrieveUserID(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	appointment, err := service.UpdateAppointment(userID, ctx.Params("id"))
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(appointment)
}

func DeleteAppointment(ctx *fiber.Ctx) error {
	userID, err := retrieveUserID(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	err = service.DeleteAppointment(userID, ctx.Params("id"))
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(model.ErrorResponse{
		Message: "Appointment deleted successfully",
	})
}

func retrieveUserID(ctx *fiber.Ctx) (string, error) {
	// retrieving token meta data
	tokenData, err := util.ExtractTokenMetadata(ctx)

	if err != nil {
		return "", err
	}

	// fetching the current logged user base on the mid retrieved from metadata
	user, err := service.GetAUser(tokenData.ID)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
