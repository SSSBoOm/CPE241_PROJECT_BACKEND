package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type maintenanceLogController struct {
	validator             domain.ValidatorUsecase
	maintenanceUsecase    domain.MaintenanceUsecase
	maintenanceLogUsecase domain.MaintenanceLogUsecase
}

func NewMaintenanceLogController(validator domain.ValidatorUsecase, maintenanceUsecase domain.MaintenanceUsecase, maintenanceLogUsecase domain.MaintenanceLogUsecase) *maintenanceLogController {
	return &maintenanceLogController{
		validator:             validator,
		maintenanceUsecase:    maintenanceUsecase,
		maintenanceLogUsecase: maintenanceLogUsecase,
	}
}

// Create godoc
// @Summary Create maintenance log
// @Description Create maintenance log
// @Tags maintenance_log
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ssid header string true "Session ID"
// @Param body body payload.MaintenanceLogCreate true "MaintenanceLogCreate"
// @Success 201 {object} domain.Response
// @Router /api/maintenance_log [post]
func (c *maintenanceLogController) Create(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.CTX_USER_ID).(string)
	var body payload.MaintenanceLogCreate
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	if maintenance, err := c.maintenanceUsecase.GetByID(body.MAINTENANCE_ID); err != nil || maintenance == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	data := &domain.MAINTENANCE_LOG{
		MAINTENANCE_ID: body.MAINTENANCE_ID,
		DESCRIPTION:    body.DESCRIPTION,
		STATUS:         body.STATUS,
		IMAGE_URL:      body.IMAGE_URL,
		STAFF_ID:       userId,
	}

	if err := c.maintenanceLogUsecase.Create(data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
