package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type promotionPriceController struct {
	validate              domain.ValidatorUsecase
	promotionPriceUsecase domain.PromotionPriceUsecase
}

func NewPromotionPriceController(validate domain.ValidatorUsecase, promotionPriceUsecase domain.PromotionPriceUsecase) *promotionPriceController {
	return &promotionPriceController{
		validate:              validate,
		promotionPriceUsecase: promotionPriceUsecase,
	}
}

// GetAll godoc
// @Summary Get all promotion prices
// @Tags promotion_price
// @Accept json
// @Produce json
// @Response 200 {object} domain.Response
// @Router /api/promotion_price [get]
func (c *promotionPriceController) GetAll(ctx *fiber.Ctx) error {
	promotionPrices, err := c.promotionPriceUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    promotionPrices,
	})
}

// GetByID godoc
// @Summary Get promotion price by ID
// @Tags promotion_price
// @Accept json
// @Produce json
// @Param id path int true "Promotion Price ID"
// @Response 200 {object} domain.Response
// @Router /api/promotion_price/{id} [get]
func (c *promotionPriceController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	promotionPrice, err := c.promotionPriceUsecase.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    promotionPrice,
	})
}
