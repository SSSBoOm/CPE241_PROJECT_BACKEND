package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type RoomTypePromotionPriceController struct {
	validator                     domain.ValidatorUsecase
	RoomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase
}

func NewRoomTypePromotionPriceController(validator domain.ValidatorUsecase, RoomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase) *RoomTypePromotionPriceController {
	return &RoomTypePromotionPriceController{
		validator:                     validator,
		RoomTypePromotionPriceUsecase: RoomTypePromotionPriceUsecase,
	}
}

// GetByRoomTypeID godoc
// @Summary Get room type promotion price by room type id
// @Description Get room type promotion price by room type id
// @Tags room_type_promotion_price
// @Accept  json
// @Produce  json
// @Param room_type_id path int true "Room Type ID"
// @Success 200 {object} []domain.ROOM_TYPE_PROMOTION_PRICE
// @Router /api/room_type_promotion_price/room_type/{room_type_id} [get]
func (c *RoomTypePromotionPriceController) GetByRoomTypeID(ctx *fiber.Ctx) error {
	roomTypeID, err := ctx.ParamsInt("room_type_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	roomTypePromotionPrices, err := c.RoomTypePromotionPriceUsecase.GetByRoomTypeID(roomTypeID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    roomTypePromotionPrices,
	})
}
