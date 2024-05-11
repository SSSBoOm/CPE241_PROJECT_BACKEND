package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentController(paymentUsecase domain.PaymentUsecase) *PaymentController {
	return &PaymentController{
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
// @Success 200 {object} domain.Response
// @Router /api/payment [post]
func (p *PaymentController) AddPaymentByUser(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	var payment payload.AddPaymentByUserIdDTO
	if err := ctx.BodyParser(&payment); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "Invalid request",
		})
	}

	paymentData := &domain.Payment{
		NAME:            payment.NAME,
		PAYMENT_NAME:    payment.PAYMENT_NAME,
		PAYMENT_NUMBER:  payment.PAYMENT_NUMBER,
		USER_ID:         id,
		PAYMENT_TYPE_ID: payment.PAYMENT_TYPE_ID,
	}

	err := p.paymentUsecase.Create(paymentData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}
