package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type PaymentTypeController struct {
	paymentTypeUsecase domain.PaymentTypeUsecase
}

func NewPaymentTypeController(paymentTypeUsecase domain.PaymentTypeUsecase) *PaymentTypeController {
	return &PaymentTypeController{
		paymentTypeUsecase: paymentTypeUsecase,
	}
}

// GetAll Payment Type	godoc
// @Summary								GetAll Payment Type
// @Description						GetAll Payment Type
// @Tags									PaymentType
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Router /api/payment_type/all [get]
func (p *PaymentTypeController) GetAll(ctx *fiber.Ctx) error {
	paymentTypes, err := p.paymentTypeUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    paymentTypes,
	})
}
