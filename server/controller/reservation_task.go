package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type ReservationTaskController struct {
	validator              domain.ValidatorUsecase
	reservationTaskUseCase domain.ReservationTaskUsecase
}

func NewReservationTaskController(validator domain.ValidatorUsecase, reservationTaskUseCase domain.ReservationTaskUsecase) *ReservationTaskController {
	return &ReservationTaskController{
		validator:              validator,
		reservationTaskUseCase: reservationTaskUseCase,
	}
}

// GetAllReservationTask godoc
// @Summary Get all reservation task
// @Description Get all reservation task
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Success 200 {object} domain.Response
// @Router /api/reservation_task [get]
func (r *ReservationTaskController) GetAllReservationTask(ctx *fiber.Ctx) error {
	data, err := r.reservationTaskUseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})
}

// CreateReservationTask godoc
// @Summary Create reservation task
// @Description Create reservation task
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body payload.ReservationTaskCreateDTO true "Create reservation task"
// @Success 201 {object} domain.Response
// @Router /api/reservation_task [post]
func (r *ReservationTaskController) CreateReservationTask(ctx *fiber.Ctx) error {
	var task payload.ReservationTaskCreateDTO
	if err := r.validator.ValidateBody(ctx, &task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	if err := r.reservationTaskUseCase.Create(&domain.RESERVATION_TASK{
		RESERVATION_ID: task.RESERVATION_ID,
		STATUS:         false,
		DATE:           task.DATE,
	}); err != nil {
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

// GetReservationTaskByReservationID godoc
// @Summary Get reservation task by reservation id
// @Description Get reservation task by reservation id
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param reservation_id path string true "Reservation ID"
// @Success 200 {object} domain.Response
// @Router /api/reservation_task/{reservation_id} [get]
func (r *ReservationTaskController) GetReservationTaskByReservationID(ctx *fiber.Ctx) error {
	param := ctx.Params("reservation_id")
	reservationID, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	data, err := r.reservationTaskUseCase.GetByReservationID(reservationID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})
}

// UpdateReservationTaskStaff godoc
// @Summary Update reservation task staff
// @Description Update reservation task staff
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body payload.ReservationTaskUpdateStaffDTO true "Update reservation task staff"
// @Success 200 {object} domain.Response
// @Router /api/reservation_task/staff [Patch]
func (r *ReservationTaskController) UpdateReservationTaskStaff(ctx *fiber.Ctx) error {
	var task payload.ReservationTaskUpdateStaffDTO
	if err := r.validator.ValidateBody(ctx, &task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	err := r.reservationTaskUseCase.UpdateStaff(task.ID, task.STAFF)
	if err != nil {
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

// UpdateReservationTaskStatus godoc
// @Summary Update reservation task status
// @Description Update reservation task status
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body payload.ReservationTaskUpdateStatusDTO true "Update reservation task status"
// @Success 200 {object} domain.Response
// @Router /api/reservation_task/status [Patch]
func (r *ReservationTaskController) UpdateReservationTaskStatus(ctx *fiber.Ctx) error {
	var task payload.ReservationTaskUpdateStatusDTO
	if err := r.validator.ValidateBody(ctx, &task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	err := r.reservationTaskUseCase.UpdateStatus(task.ID, task.STATUS)
	if err != nil {
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
