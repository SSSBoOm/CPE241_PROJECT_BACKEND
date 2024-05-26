package controller

import (
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type promotionPriceController struct {
	validate                      domain.ValidatorUsecase
	promotionPriceUsecase         domain.PromotionPriceUsecase
	roomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase
}

func NewPromotionPriceController(validate domain.ValidatorUsecase, promotionPriceUsecase domain.PromotionPriceUsecase, roomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase) *promotionPriceController {
	return &promotionPriceController{
		validate:                      validate,
		promotionPriceUsecase:         promotionPriceUsecase,
		roomTypePromotionPriceUsecase: roomTypePromotionPriceUsecase,
	}
}

// GetAll godoc
// @Summary Get all promotion prices
// @Tags promotion_price
// @Accept json
// @Produce json
// @Response 200 {object} []domain.PROMOTION_PRICE
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
// @Response 200 {object} domain.PROMOTION_PRICE
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

// Create godoc
// @Summary Create promotion price
// @Tags promotion_price
// @Accept json
// @Produce json
// @Param promotion_price body payload.PromotionPriceCreateDTO true "Promotion Price"
// @Response 201 {object} domain.PROMOTION_PRICE
// @Router /api/promotion_price [post]
func (c *promotionPriceController) Create(ctx *fiber.Ctx) error {
	var body payload.PromotionPriceCreateDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		fmt.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	promotionPrice := &domain.PROMOTION_PRICE{
		NAME:       body.NAME,
		PRICE:      body.PRICE,
		START_DATE: body.START_DATE,
		END_DATE:   body.END_DATE,
	}

	id, err := c.promotionPriceUsecase.Create(promotionPrice)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	for _, roomTypeId := range body.ROOM_TYPE_ID {
		roomTypePromotionPrice := &domain.ROOM_TYPE_PROMOTION_PRICE{
			ROOM_TYPE_ID:       roomTypeId,
			PROMOTION_PRICE_ID: *id,
			IS_ACTIVE:          true,
		}
		_, err := c.roomTypePromotionPriceUsecase.Create(roomTypePromotionPrice)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			})
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    promotionPrice,
	})
}
