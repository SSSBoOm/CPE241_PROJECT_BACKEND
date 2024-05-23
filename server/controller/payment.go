package controller

import (
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	validator      domain.ValidatorUsecase
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentController(validator domain.ValidatorUsecase, paymentUsecase domain.PaymentUsecase) *PaymentController {
	return &PaymentController{
		validator:      validator,
		paymentUsecase: paymentUsecase,
	}
}

// AddPaymentByUser godoc
// @Summary Add payment by user
// @Description Add payment by user
// @Tags payment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body payload.AddPaymentByUserIdDTO true "Add payment by user"
// @Success 201 {object} domain.Response
// @Router /api/payment [post]
func (c *PaymentController) AddPaymentByUser(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	var body payload.AddPaymentByUserIdDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "Invalid request",
		})
	}

	paymentData := &domain.Payment{
		NAME:               body.NAME,
		PAYMENT_FIRST_NAME: body.PAYMENT_FIRST_NAME,
		PAYMENT_LAST_NAME:  body.PAYMENT_LAST_NAME,
		PAYMENT_NUMBER:     body.PAYMENT_NUMBER,
		USER_ID:            id,
		PAYMENT_TYPE_ID:    body.PAYMENT_TYPE_ID,
	}

	err := c.paymentUsecase.Create(paymentData)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}
