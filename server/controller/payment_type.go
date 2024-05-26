package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"

	"github.com/gofiber/fiber/v2"
)

type PaymentTypeController struct {
	validator          domain.ValidatorUsecase
	paymentTypeUsecase domain.PaymentTypeUsecase
}

func NewPaymentTypeController(validator domain.ValidatorUsecase, paymentTypeUsecase domain.PaymentTypeUsecase) *PaymentTypeController {
	return &PaymentTypeController{
		validator:          validator,
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
// @Response							200 {object} []domain.PaymentType
// @Router /api/payment_type [get]
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

// GetByID Payment Type	godoc
// @Summary								GetByID Payment Type
// @Description						GetByID Payment Type
// @Tags									PaymentType
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									id path int true "Payment Type ID"
// @Response							200 {object} domain.PaymentType
// @Router /api/payment_type/{id} [get]
func (p *PaymentTypeController) GetByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}
	paymentType, err := p.paymentTypeUsecase.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    paymentType,
	})
}

// Create Payment Type	godoc
// @Summary								Create Payment Type
// @Description						Create Payment Type
// @Tags									PaymentType
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									name body payload.PaymentTypeCreateDTO true "Payment Type Name"
// @Response							200 {object} domain.Response
// @Router /api/payment_type [post]
func (c *PaymentTypeController) Create(ctx *fiber.Ctx) error {
	var body payload.PaymentTypeCreateDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	paymentType := &domain.PaymentType{
		NAME: body.NAME,
	}

	if err := c.paymentTypeUsecase.Create(paymentType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// Update Payment Type	godoc
// @Summary								Update Payment Type
// @Description						Update Payment Type
// @Tags									PaymentType
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									id path int true "Payment Type ID"
// @Param									name body payload.PaymentTypeUpdateDTO true "Payment Type Name"
// @Response							200 {object} domain.Response
// @Router /api/payment_type [put]
func (c *PaymentTypeController) Update(ctx *fiber.Ctx) error {
	var body payload.PaymentTypeUpdateDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	paymentType := &domain.PaymentType{
		ID:   body.ID,
		NAME: body.NAME,
	}

	if err := c.paymentTypeUsecase.Update(paymentType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
