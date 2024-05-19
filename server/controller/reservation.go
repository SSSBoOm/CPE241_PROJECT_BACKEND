package controller

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type reservationController struct {
	validator          domain.ValidatorUsecase
	reservationUsecase domain.ReservationUsecase
	roomUsecase        domain.RoomUsecase
	roomTypeUsecase    domain.RoomTypeUsecase
}

func NewReservationController(validator domain.ValidatorUsecase, reservationUsecase domain.ReservationUsecase, roomUsecase domain.RoomUsecase, roomTypeUsecase domain.RoomTypeUsecase) *reservationController {
	return &reservationController{
		validator:          validator,
		reservationUsecase: reservationUsecase,
		roomUsecase:        roomUsecase,
		roomTypeUsecase:    roomTypeUsecase,
	}
}

// createReservation godoc
// @Summary Create reservation
// @Description Create reservation
// @Tags reservation
// @Accept json
// @Produce json
// @Param body body payload.CreateReservationDTO true "Create reservation"
// @Success 200 {object} domain.Response
// @Router /api/reservation [post]
func (c *reservationController) CreateReservation(ctx *fiber.Ctx) error {
	userID := ctx.Locals(constant.CTX_USER_ID).(string)
	var body payload.CreateReservationDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	if _, err := c.reservationUsecase.Create(&domain.RESERVATION{
		ROOM_ID:         &body.ROOM_ID,
		USER_ID:         userID,
		START_DATE:      body.START_DATE,
		END_DATE:        body.END_DATE,
		PRICE:           body.PRICE,
		STATUS:          domain.RESERVATION_STATUS_WAITING_APPROVE_PAYMENT,
		PAYMENT_DATE:    time.Now(),
		PAYMENT_INFO_ID: body.PAYMENT_INFO_ID,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// getReservationByUserID godoc
// @Summary Get reservation by user id
// @Description Get reservation by user id
// @Tags reservation
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} domain.Response
// @Router /api/reservation/me/all [get]
func (c *reservationController) GetReservationByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Locals(constant.CTX_USER_ID).(string)

	data, err := c.reservationUsecase.GetByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})

}

// getRoomAvailableGroupByRoomType godoc
// @Summary Get room available group by room type
// @Description Get room available group by room type
// @Tags reservation
// @Accept json
// @Produce json
// @Param body body payload.GetRoomAvailableGroupByRoomTypeDTO true "Get room available group by room type"
// @Success 200 {object} domain.Response
// @Router /api/reservation/get-room-available [post]
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
