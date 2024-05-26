package controller

import (
	"fmt"
	"strconv"
	"sync"
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

// getAll godoc
// @Summary Get all reservation
// @Description Get all reservation
// @Tags reservation
// @Accept json
// @Produce json
// @Success 200 {object} []domain.RESERVATION
// @Router /api/reservation [get]
func (c *reservationController) GetAll(ctx *fiber.Ctx) error {
	data, err := c.reservationUsecase.GetAll()
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

// GetByID godoc
// @Summary Get reservation by id
// @Description Get reservation by id
// @Tags reservation
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} domain.RESERVATION
// @Router /api/reservation/{id} [get]
func (c *reservationController) GetByID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}
	data, err := c.reservationUsecase.GetByID(id)
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
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	if (body.ROOM_ID == nil && body.SERVICE_ID == nil) || (body.ROOM_ID != nil && body.SERVICE_ID != nil) {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	if _, err := c.reservationUsecase.Create(&domain.RESERVATION{
		TYPE:              body.TYPE,
		ROOM_ID:           body.ROOM_ID,
		SERVICE_ID:        body.SERVICE_ID,
		USER_ID:           userID,
		START_DATE:        body.START_DATE,
		END_DATE:          body.END_DATE,
		PRICE:             *body.PRICE,
		ROOM_PROMOTION_ID: body.ROOM_PROMOTION_ID,
		STATUS:            domain.RESERVATION_STATUS_WAITING_APPROVE_PAYMENT,
		PAYMENT_DATE:      time.Now(),
		PAYMENT_INFO_ID:   body.PAYMENT_INFO_ID,
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
// @Success 200 {object} []domain.RESERVATION
// @Router /api/reservation/me [get]
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

// getReservationByReservationType godoc
// @Summary Get reservation by reservation type
// @Description Get reservation by reservation type
// @Tags reservation
// @Accept json
// @Produce json
// @Param type path string true "Reservation Type"
// @Success 200 {object} []domain.RESERVATION
// @Router /api/reservation/type/{type} [get]
func (c *reservationController) GetReservationByReservationType(ctx *fiber.Ctx) error {
	param := ctx.Params("type")
	reservationType := domain.RESERVATION_TYPE(param)
	data, err := c.reservationUsecase.GetByType(reservationType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	var wg sync.WaitGroup
	wg.Add(len(*data))
	for i, item := range *data {
		go func(i int, item domain.RESERVATION) {
			defer wg.Done()
			(*data)[i].USER_ID = ""
			(*data)[i].USER = nil
			(*data)[i].PAYMENT_INFO = nil
		}(i, item)
	}
	wg.Wait()

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})
}

// UpdateStaff godoc
// @Summary Update staff
// @Description Update staff
// @Tags reservation
// @Accept json
// @Produce json
// @Param body body payload.UpdateReservationStaffDTO true "Update staff"
// @Success 200 {object} domain.Response
// @Router /api/reservation/staff [patch]
func (c *reservationController) UpdateStaff(ctx *fiber.Ctx) error {
	var body payload.UpdateReservationStaffDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}
	c.reservationUsecase.UpdateStaff(body.RESERVATION_ID, body.STAFF_ID)

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// updateStatus godoc
// @Summary Update status
// @Description Update status
// @Tags reservation
// @Accept json
// @Produce json
// @Param body body payload.UpdateReservationStatusDTO true "Update status"
// @Success 200 {object} domain.Response
// @Router /api/reservation/status [patch]
func (c *reservationController) UpdateStatus(ctx *fiber.Ctx) error {
	var body payload.UpdateReservationStatusDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}
	c.reservationUsecase.UpdateStatus(body.RESERVATION_ID, body.STATUS)

	return ctx.JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// updatePayment godoc
// @Summary Update payment
// @Description Update payment
// @Tags reservation
// @Accept json
// @Produce json
// @Param body body payload.UpdateReservationPaymentDTO true "Update payment"
// @Success 200 {object} domain.Response
// @Router /api/reservation/payment [patch]
func (c *reservationController) UpdatePayment(ctx *fiber.Ctx) error {
	var body payload.UpdateReservationPaymentDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}
	if err := c.reservationUsecase.UpdatePayment(body.RESERVATION_ID, body.PAYMENT_INFO_ID); err != nil {
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
