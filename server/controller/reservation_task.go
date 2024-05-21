package controller

import (
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
