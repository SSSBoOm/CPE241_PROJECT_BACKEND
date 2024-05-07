package controller

import (
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase    domain.UserUsecase
	paymentUsecase domain.PaymentUsecase
}

func NewUserController(userUsecase domain.UserUsecase, paymentUsecase domain.PaymentUsecase) *UserController {
	return &UserController{
		userUsecase:    userUsecase,
		paymentUsecase: paymentUsecase,
	}
}

func (u *UserController) Me(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	user, err := u.userUsecase.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "User not found",
		})
	}
	user.ID = ""

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    user,
	})
}

func (u *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	fmt.Println(id)
	user, err := u.userUsecase.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_USER_NOT_FOUND,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    user,
	})
}

func (u *UserController) UpdateByID(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}

func (u *UserController) UpdateRoleByID(ctx *fiber.Ctx) error {
	var body payload.UpdateUserRoleDTO
	err := validator.NewPayloadValidator().ValidateBody(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	err = u.userUsecase.UpdateRoleById(body.USER_ID, body.ROLE_ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}

// UpdateUserInformation	godoc
// @Summary								Update user information
// @Description						Update user information
// @Tags									user
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									payload body	payload.UpdateUserDTO true "Payload"
// @Router /api/user/			[patch]
func (u *UserController) UpdateInfomationByID(ctx *fiber.Ctx) error {
	var body payload.UpdateUserDTO
	err := validator.NewPayloadValidator().ValidateBody(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	user := &domain.User{
		ID:         ctx.Locals(constant.CTX_USER_ID).(string),
		PREFIX:     body.PREFIX,
		FIRST_NAME: body.FIRST_NAME,
		LAST_NAME:  body.LAST_NAME,
		DOB:        body.DOB,
		PHONE:      body.PHONE,
		GENDER:     body.GENDER,
		ADDRESS:    body.ADDRESS,
	}

	err = u.userUsecase.UpdateInfomation(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}

func (u *UserController) GetPaymentByUserID(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	payments, err := u.paymentUsecase.GetByUserID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    payments,
	})
}
