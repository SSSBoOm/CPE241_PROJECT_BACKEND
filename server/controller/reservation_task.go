package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type ReservationTaskController struct {
	reservationTaskUseCase domain.ReservationTaskUsecase
}

func NewReservationTaskController(reservationTaskUseCase domain.ReservationTaskUsecase) *ReservationTaskController {
	return &ReservationTaskController{
		reservationTaskUseCase: reservationTaskUseCase,
	}
}

// CreateReservationTask godoc
// @Summary Create reservation task
// @Description Create reservation task
// @Tags reservation_task
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body domain.RESERVATION_TASK true "Create reservation task"
// @Success 200 {object} domain.Response
// @Router /api/reservation_task [post]
func (r *ReservationTaskController) CreateReservationTask(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	var task domain.RESERVATION_TASK
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: "Invalid request",
		})
	}

	task.STAFF_ID = &id
	err := r.reservationTaskUseCase.Create(&domain.RESERVATION_TASK{
		RESERVATION_ID: task.RESERVATION_ID,
	})
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
			MESSAGE: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    data,
	})
}
