package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (u *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := u.userUsecase.FindById(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "User not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    user,
	})
}

func (u *UserController) UpdateByID(ctx *fiber.Ctx) error {
	// var body payload.UpdateUser
	// err := validator.ValidateBody(c, &body)
	// if err != nil {
	// 	return response.ErrorResponse(c, 400, err.Error())
	// }

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}
