package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type reservationController struct {
	validator       domain.ValidatorUsecase
	roomUsecase     domain.RoomUsecase
	roomTypeUsecase domain.RoomTypeUsecase
}

func NewReservationController(validator domain.ValidatorUsecase, roomUsecase domain.RoomUsecase, roomTypeUsecase domain.RoomTypeUsecase) *reservationController {
	return &reservationController{
		validator:       validator,
		roomUsecase:     roomUsecase,
		roomTypeUsecase: roomTypeUsecase,
	}
}

func (c *reservationController) GetRoomAvailableGroupByRoomType(ctx *fiber.Ctx) error {
	var body payload.GetRoomAvailableGroupByRoomTypeDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
