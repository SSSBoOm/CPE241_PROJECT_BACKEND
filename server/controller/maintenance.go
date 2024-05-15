package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type maintenanceController struct {
	maintenanceUsecase domain.MaintenanceUsecase
}

func NewMaintenanceController(maintenanceUsecase domain.MaintenanceUsecase) *maintenanceController {
	return &maintenanceController{
		maintenanceUsecase: maintenanceUsecase,
	}
}

// GetAll Maintenance	godoc
// @Summary								GetAll Maintenance
// @Description						GetAll Maintenance
// @Tags									Maintenance
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Router /api/maintenance/all [get]
func (c *maintenanceController) GetAll(ctx *fiber.Ctx) error {
	maintenances, err := c.maintenanceUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    maintenances,
	})
}
